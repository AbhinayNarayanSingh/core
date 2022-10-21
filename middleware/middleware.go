package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "users")

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Authentication key not provided"})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Authentication key invalid."})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.M{"_id": claims.ID, "email": claims.Email, "phone": claims.Phone, "isactive": claims.IsAdmin, "isadmin": claims.IsAdmin}

		if count, error := userCollection.CountDocuments(ctx, filter); error != nil && count != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Authentication key invalid."})
			c.Abort()
			return
		}

		c.Set("_id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("is_admin", claims.IsAdmin)
		c.Set("is_active", claims.IsActive)
		c.Next()
	}
}

func AdminAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthenticationMiddleware()
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Admin authentication key required."})
			c.Abort()
			return
		}
	}
}
