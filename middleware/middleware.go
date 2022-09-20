package middleware

import (
	"net/http"

	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
)

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
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
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
