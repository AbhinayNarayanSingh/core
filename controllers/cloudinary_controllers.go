package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

type payloadStruct struct {
	Public_id *string `json:"public_id,omitempty" bson:"public_id,omitempty" validate:"required"`
}

func CloudinaryDestroy() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload *payloadStruct

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.BadRequest, "details": err.Error()})
			return
		}

		// Create a new Cloudinary uploader
		CLOUDINARY_CLOUD_NAME := os.Getenv("CLOUDINARY_CLOUD_NAME")
		CLOUDINARY_API_KEY := os.Getenv("CLOUDINARY_API_KEY")
		CLOUDINARY_API_SECRET := os.Getenv("CLOUDINARY_API_SECRET")

		cloudinaryURL := fmt.Sprintf("cloudinary://%s:%s@%s", CLOUDINARY_API_KEY, CLOUDINARY_API_SECRET, CLOUDINARY_CLOUD_NAME)

		cld, err := cloudinary.NewFromURL(cloudinaryURL)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.BadRequest, "details": err.Error()})
			return
		}

		deleteResult, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
			PublicID: *payload.Public_id,
		})

		if err != nil {
			c.JSON(404, gin.H{"message": "Failed to delete image from Cloudinary", "publicId": payload.Public_id, "error": err})
			return
		}

		c.JSON(200, deleteResult)
	}
}
