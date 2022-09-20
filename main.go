package main

import (
	"os"

	"github.com/AbhinayNarayanSingh/core/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Helper function to return a URL pattern
	routes.Path(router)
	routes.SecurePath(router)

	router.Run(":" + port)
}
