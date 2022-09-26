package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = config.OpenCollection(config.Client, "products")
var productPriceCollection *mongo.Collection = config.OpenCollection(config.Client, "product_prices")
var productImageCollection *mongo.Collection = config.OpenCollection(config.Client, "product_images")

// create product, productPrice, productImage
func ProductCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.ProductPayload

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err.Error()})
			return
		}

		Product_UID := primitive.NewObjectID()

		payload.ProductDetail.ID = Product_UID

		for _, item := range *payload.Product_Price {
			var newProductPrice models.ProductPrice
			Price_UID := primitive.NewObjectID()

			newProductPrice = item
			newProductPrice.ID = Price_UID
			newProductPrice.Product_ID = Product_UID

			if _, err := productPriceCollection.InsertOne(ctx, newProductPrice); err != nil {
				fmt.Println("error during insertion of product price")
			}
		}

		for _, item := range *payload.Product_Images {
			var newProductImage models.ProductImage

			Image_UID := primitive.NewObjectID()

			newProductImage = item
			newProductImage.ID = Image_UID
			newProductImage.Product_ID = Product_UID

			if _, err := productImageCollection.InsertOne(ctx, newProductImage); err != nil {
				fmt.Println("error during insertion of product image")
			}
		}

		if _, err := productCollection.InsertOne(ctx, payload.ProductDetail); err != nil {
			fmt.Println("User data not created")
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
			return
		}

		c.JSON(200, gin.H{"message": Product_UID})
	}
}

// delete product, productPrice, productImage
func ProductDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.ProductPayload

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err.Error()})
			return
		}

		product_id, _ := primitive.ObjectIDFromHex(*payload.Product_ID)
		filter := bson.M{"_id": product_id}

		if _, err := productCollection.DeleteMany(ctx, filter); err != nil {
			fmt.Println("error occurs during productCollection.DeleteMany")
		}

		filter = bson.M{"product_id": product_id}

		if _, err := productPriceCollection.DeleteMany(ctx, filter); err != nil {
			fmt.Println("error occurs during productPriceCollection.DeleteMany")
		}

		if _, err := productImageCollection.DeleteMany(ctx, filter); err != nil {
			fmt.Println("error occurs during productImageCollection.DeleteMany")
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product Deleted"})
	}
}

// get product_images
func ProductGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.ProductPayload

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err.Error()})
			return
		}

		product_uid, _ := primitive.ObjectIDFromHex(*payload.Product_ID)

		var GetProductDetailsFn = func() {
			filter := bson.M{"_id": product_uid}
			var result bson.M
			if err := productCollection.FindOne(ctx, filter).Decode(&result); err != nil {
				c.JSON(400, gin.H{"message": err})
			}
			c.JSON(200, gin.H{"response": result})
		}

		var GetProductImagesFn = func() {
			filter := bson.M{"product_id": product_uid, "variant": payload.Variant_Color}

			var result bson.M
			if err := productImageCollection.FindOne(ctx, filter).Decode(&result); err != nil {
				c.JSON(400, gin.H{"message": err})
				return
			}
			c.JSON(200, gin.H{"response": result})
		}

		var GetProductPriceFn = func() {
			filter := bson.M{"product_id": product_uid, "size": payload.Variant_Size}

			var result bson.M
			if err := productPriceCollection.FindOne(ctx, filter).Decode(&result); err != nil {
				c.JSON(400, gin.H{"message": err})
				return
			}
			c.JSON(200, gin.H{"response": result})
		}

		switch payload.Operation {
		case 1:
			GetProductDetailsFn()
		case 2:
			GetProductImagesFn()
		case 3:
			GetProductPriceFn()
		}
	}
}

// get product_price

// get product_details

// update product

// update price
