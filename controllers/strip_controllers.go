package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
	"github.com/stripe/stripe-go/v74/webhook"
)

func CreatePaymentIntent() gin.HandlerFunc {
	return func(c *gin.Context) {
		stripe.Key = os.Getenv("STRIP_KEY")

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(100050),
			Currency: stripe.String(string(stripe.CurrencyINR)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}
		pi, err := paymentintent.New(params)
		// pi, err := paymentintent.Get("payment_intent_id", nil)

		if err != nil {
			// Try to safely cast a generic error to a stripe.Error so that we can get at
			// some additional Stripe-specific information about what went wrong.
			if stripeErr, ok := err.(*stripe.Error); ok {
				c.JSON(400, stripeErr.Error())
			} else {
				c.JSON(500, err.Error())
			}
			return
		}

		c.JSON(http.StatusSeeOther, struct {
			ClientSecret string `json:"clientSecret"`
		}{
			ClientSecret: pi.ClientSecret,
		})

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
