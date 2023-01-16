package main

import (
	"io"
	"log"
	"os"

	"github.com/AbhinayNarayanSingh/core/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	if file, err := os.OpenFile("./logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(file)
	}

	if logfile, err := os.OpenFile("./logs/httprequest.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		log.Fatal(err)
	} else {
		gin.DefaultWriter = io.MultiWriter(logfile)
	}

	gin.SetMode(gin.ReleaseMode)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Helper function to return a URL pattern

	// router.StaticFile("/", "static")
	// routes.WebsocketPath(router)

	routes.Path(router)
	routes.AdminSecurePath(router)
	routes.SecurePath(router)

	router.Run(":" + port)
}
