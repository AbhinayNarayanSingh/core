package controllers

import (
	"context"
	"time"

	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Welcome() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(200, gin.H{"message": "Hello programmer..."})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": locals.InternalServerError})
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		userId := c.Param("user_id")
		uid, _ := primitive.ObjectIDFromHex(userId)
		if err := utils.AuthenticateUser(c, userId); err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := userCollection.FindOne(ctx, bson.M{"_id": uid}).Decode(&user); err != nil {
			c.JSON(500, gin.H{"message": "User not found.", "details": err.Error()})
			return
		}
		c.JSON(200, user)
	}
}

// make & remove admin

// delete user
