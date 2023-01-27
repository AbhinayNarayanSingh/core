package routes

import (
	"github.com/AbhinayNarayanSingh/core/controllers"
	"github.com/AbhinayNarayanSingh/core/middleware"
	"github.com/gin-gonic/gin"
)

func WebsocketPath(router *gin.Engine) {
	router.GET("/ws", controllers.ChatController())
}

func Path(router *gin.Engine) {

	router.POST("/signup", controllers.SignUp())
	router.POST("/signin", controllers.SignIn())
	router.POST("/signin/otp", controllers.OTPVerification(4))

	router.POST("/otp", controllers.OTPVerificationInitiator(0))
	router.POST("/user/password/reset", controllers.OTPVerification(6))

	router.POST("/otp/verify", controllers.OTPVerification(0))

	router.POST("/user/password/reset/verify", controllers.OTPVerificationInitiator(6))

	router.POST("/product", controllers.ProductGet())
	router.POST("/address", controllers.SaveAddress())
}

func SecurePath(router *gin.Engine) {
	router.Use(middleware.AuthenticationMiddleware())

	router.POST("/user/password", controllers.UpdatePassword())
	router.GET("/users/:user_id", controllers.GetUserByID())
}

func AdminSecurePath(router *gin.Engine) {
	// router.Use(middleware.AdminAuthenticationMiddleware())

	router.GET("/users", controllers.GetUsers())

	router.POST("/products", controllers.ProductCreate())
	router.PATCH("/products", controllers.ProductUpdate())
	router.DELETE("/products", controllers.ProductDelete())
}
