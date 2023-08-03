package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
	"github.com/stripe/stripe-go/v74/webhook"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionsCollection *mongo.Collection = config.OpenCollection(config.Client, "transactions")
var ServicesCollection *mongo.Collection = config.OpenCollection(config.Client, "services")

var stringempty string = ""

func findAllServices(serviceIDs []primitive.ObjectID) ([]models.Services, error) {
	var services []models.Services
	filter := bson.M{"_id": bson.M{"$in": serviceIDs}}
	cur, err := ServicesCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var service models.Services
		err := cur.Decode(&service)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

func CreatePaymentIntent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Stripe
		stripe.Key = os.Getenv("STRIP_KEY")

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.BadRequest, "details": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		serviceIDs := make([]primitive.ObjectID, len(payload.Services))
		for i, item := range payload.Services {
			serviceID, err := primitive.ObjectIDFromHex(item.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
				return
			}
			serviceIDs[i] = serviceID
		}

		services, err := findAllServices(serviceIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch services"})
			return
		}

		// map to associate service IDs with their details
		serviceMap := make(map[string]models.Services)
		for _, service := range services {
			serviceMap[service.ID.Hex()] = service
		}

		totalCost := float32(0)

		for _, item := range payload.Services {
			service, found := serviceMap[item.ID]
			if !found {
				c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
				return
			}
			cost := service.BasePrice * float32(item.Duration)
			totalCost += cost
		}

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(int64(totalCost * 100)),
			Currency: stripe.String(string(stripe.CurrencyCAD)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}
		pi, err := paymentintent.New(params)

		if err != nil {
			if stripeErr, ok := err.(*stripe.Error); ok {
				c.JSON(500, stripeErr.Error())
			} else {
				c.JSON(500, err.Error())
			}
			return
		}

		payload.ID = primitive.NewObjectID()
		payload.Against, _ = primitive.ObjectIDFromHex(*payload.Listing_ID)
		payload.User_ID, _ = primitive.ObjectIDFromHex(*payload.Seller_ID)

		payload.ClientSecret = pi.ClientSecret
		payload.PaymentIntent_ID = pi.ID
		payload.Status = string(pi.Status)

		if _, err := transactionsCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"_id":             payload.ID,
			"clientSecret":    payload.ClientSecret,
			"paymentIntentID": payload.PaymentIntent_ID,
			"totalCost":       totalCost,
			"type":            payload.Type,
			"narration":       payload.Narration,
		})
	}
}

func CancelPaymentIntent() gin.HandlerFunc {
	return func(c *gin.Context) {
		stripe.Key = os.Getenv("STRIP_KEY")
		paymentIntentID := "pi_3NZvLaSAIDSoJ9VZ0CvysG9F"
		pi, err := paymentintent.Cancel(paymentIntentID, nil)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(http.StatusOK, pi.Status)
	}
}

func StatusPaymentIntent() gin.HandlerFunc {
	return func(c *gin.Context) {
		stripe.Key = os.Getenv("STRIP_KEY")
		paymentIntentID := "pi_3NFDkXSAIDSoJ9VZ1lTrKlG8"
		pi, err := paymentintent.Get(paymentIntentID, nil)
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(http.StatusOK, pi.Status)
	}
}

func StripeWebhookListener() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the request body
		payload, err := c.GetRawData()
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Verify the signature
		event, err := webhook.ConstructEvent(payload, c.Request.Header.Get("Stripe-Signature"), "YOUR_STRIPE_WEBHOOK_SECRET")
		if err != nil {
			log.Printf("Failed to verify webhook signature: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Handle the event based on its type
		switch event.Type {
		case "payment_intent.succeeded":
			// Handle payment intent succeeded event
			log.Println("Payment intent succeeded")
			// Additional handling logic here
		case "payment_intent.payment_failed":
			// Handle payment intent failed event
			log.Println("Payment intent failed")
			// Additional handling logic here
		default:
			// Ignore unrecognized event types
			log.Println("Unrecognized event type")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Return a success response
		c.JSON(http.StatusOK, gin.H{"message": "done", "event": event})

	}

}

func CreateCheckoutSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := os.Getenv("FRONTEND_DOMAIN")

		stripe.Key = os.Getenv("STRIP_KEY")

		// Create a product
		productParams := &stripe.ProductParams{
			Name:        stripe.String("Paid Listing Service"), // You can add more details to the product if needed
			Description: stripe.String("Paid Listing - Visibility for 2 Weeks in Crampton, Ontario, Canada"),
		}
		product, _ := product.New(productParams)

		// Create a price
		priceParams := &stripe.PriceParams{
			Product:    stripe.String(product.ID),
			UnitAmount: stripe.Int64(1000),
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
		}
		price, _ := price.New(priceParams)

		params := &stripe.CheckoutSessionParams{
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					Price:    stripe.String(price.ID),
					Quantity: stripe.Int64(1),
				},
			},

			Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
			SuccessURL: stripe.String(domain + "/success.html"),
			CancelURL:  stripe.String(domain + "/cancel.html"),
		}

		s, err := session.New(params)

		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(http.StatusSeeOther, gin.H{"message": "done", "url": s.URL, "product": product})

	}
}
