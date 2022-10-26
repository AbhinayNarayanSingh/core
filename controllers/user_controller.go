package controllers

import (
	"context"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var addressCollection *mongo.Collection = config.OpenCollection(config.Client, "address")

func SaveAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.Address

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.BadRequest})
			return
		}

		UID := primitive.NewObjectID()
		payload.ID = UID
		payload.Created_at = utils.TimeStampFn()

		if _, err := addressCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(500, gin.H{"message": "failed during the insertion of address"})
			return
		}

		c.JSON(200, gin.H{"message": payload})
	}
}
