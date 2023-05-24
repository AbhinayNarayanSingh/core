package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var listingCollection *mongo.Collection = config.OpenCollection(config.Client, "listings")
var categoriesCollection *mongo.Collection = config.OpenCollection(config.Client, "categories")

func CreateNewCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Category

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.BadRequest, "details": err.Error()})
			return
		}
		object_id := primitive.NewObjectID()
		payload.ID = object_id

		if _, err := categoriesCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
			return
		}

		c.JSON(201, gin.H{"message": "done"})
	}
}

func GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []bson.M

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := categoriesCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve categories", "error": err.Error()})
			return
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var category bson.M
			if err := cursor.Decode(&category); err != nil {
				log.Fatal(err)
			}
			categories = append(categories, category)
		}

		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred during category retrieval", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"categories": categories})
	}
}

func CreateNewListing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Listing

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.BadRequest, "details": err.Error()})
			return
		}

		object_id := primitive.NewObjectID()
		timestamp := utils.TimeStampFn()

		payload.ID = object_id
		payload.Posted_on = timestamp
		payload.Updated_on = timestamp

		if _, err := listingCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
			return
		}

		c.JSON(201, gin.H{"message": "done"})
	}
}
