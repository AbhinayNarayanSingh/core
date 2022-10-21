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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productCollection *mongo.Collection = config.OpenCollection(config.Client, "products")
var productVarientsCollection *mongo.Collection = config.OpenCollection(config.Client, "product_varients")
var productImageCollection *mongo.Collection = config.OpenCollection(config.Client, "product_images")

// create product
func ProductCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.ProductPayload

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.InternalServerError, "details": err.Error()})
			return
		}

		Product_UID := primitive.NewObjectID()
		payload.ProductDetail.ID = Product_UID
		payload.Varients.Product_ID = Product_UID

		if _, err := productVarientsCollection.InsertOne(ctx, payload.Varients); err != nil {
			fmt.Println("error during insertion of product varients")
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

		c.JSON(200, gin.H{"message": payload})
	}
}

// delete product
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

		if _, err := productVarientsCollection.DeleteMany(ctx, filter); err != nil {
			fmt.Println("error occurs during productVarientsCollection.DeleteMany")
		}

		if _, err := productImageCollection.DeleteMany(ctx, filter); err != nil {
			fmt.Println("error occurs during productImageCollection.DeleteMany")
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product Deleted"})
	}
}

// update product
func ProductUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.ProductPayload

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.InternalServerError})
		}

		Product_UID, _ := primitive.ObjectIDFromHex(*payload.Product_ID)

		var UpdateProductFn = func(collection *mongo.Collection, update *bson.D) {
			upsert := true
			filter := bson.M{"product_id": Product_UID}
			opts := options.UpdateOptions{Upsert: &upsert}

			if _, err := collection.UpdateOne(ctx, filter, update, &opts); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": locals.InternalServerError})
			}

			c.JSON(http.StatusOK, gin.H{"message": "Done"})
		}

		switch payload.Operation {
		case 1:
			update := bson.D{{Key: "$set", Value: payload.ProductDetail}}
			UpdateProductFn(productCollection, &update)
		case 2:
			update := bson.D{{Key: "$set", Value: payload.Varients}}
			UpdateProductFn(productVarientsCollection, &update)
		case 3:
			update := bson.D{{Key: "$set", Value: payload.Product_Images}}
			UpdateProductFn(productImageCollection, &update)
		}

	}
}

// get product
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

		var GetProductsFn = func(collection *mongo.Collection, filter *bson.M) {
			var result bson.M
			if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
				c.JSON(400, gin.H{"message": err})
				return
			}
			c.JSON(200, gin.H{"response": result})
		}

		switch payload.Operation {
		case 1:
			filter := bson.M{"_id": product_uid}
			GetProductsFn(productCollection, &filter)
		case 2:
			filter := bson.M{"product_id": product_uid}
			GetProductsFn(productVarientsCollection, &filter)
		case 3:
			filter := bson.M{"product_id": product_uid, "variant": payload.Variant_Color}
			GetProductsFn(productImageCollection, &filter)
		}
	}
}
