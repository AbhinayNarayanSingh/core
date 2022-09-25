package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = config.OpenCollection(config.Client, "products")

func ProductCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Product

		// var newProduct models.Product
		// var newProductPrice models.ProductPrice
		// var newProductImage models.ProductImage

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err.Error()})
			return
		}
		payload.ID = primitive.NewObjectID()
		paylodId := payload.ID.Hex()

		payload.Product_Id = &paylodId

		if _, err := productCollection.InsertOne(ctx, payload); err != nil {
			fmt.Println("User data not created")
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
			return
		}
		c.JSON(200, gin.H{"message": payload})
	}

}

// create product, productPrice, productImage

// delete product, productPrice, productImage

// update product

// update price

// get product_images from productImageCollection

// get product_price from productPriceColletion

// get product_details
