package routes

import (
	"github.com/AbhinayNarayanSingh/core/controllers"
	"github.com/gin-gonic/gin"
)

func WebsocketPath(router *gin.Engine) {
	// router.GET("/ws", controllers.ChatController())
}

func Path(router *gin.Engine) {

	router.GET("/", controllers.Welcome())

	router.POST("/otp", controllers.OTPVerificationInitiator(0)) // one route to generate otp for all pourpose

	// user signup routes
	router.POST("/signup", controllers.SignUp())
	router.POST("/signup/verify", controllers.OTPVerification(4))

	// user authentication routes
	router.POST("/signin", controllers.SignIn())
	router.POST("/refresh-token", controllers.GetUserByToken())

	// password reset routes
	router.POST("/password-reset", controllers.OTPVerification(5)) // 5.email password reset		6. phone password reset
	router.POST("/user/verify", controllers.OTPVerification(1))    // 1.email verification		2.phone number verification

	router.POST("/address", controllers.SaveAddress())

	router.POST("/listing", controllers.CreateNewListing())
	router.GET("/listings", controllers.GetListings())

	// categories
	router.POST("/categories", controllers.CreateNewCategory())
	router.GET("/categories", controllers.GetCategories())

	router.POST("/services", controllers.CreateNewService())
	router.GET("/services", controllers.GetServices())

	router.DELETE("/cloudinary/destroy", controllers.CloudinaryDestroy())

}

func SecurePath(router *gin.Engine) {
	// router.Use(middleware.AuthenticationMiddleware())

	router.POST("/user/password", controllers.UpdatePassword())
	router.GET("/users/:user_id", controllers.GetUserByID())

	router.POST("/create-checkout-session", controllers.CreateCheckoutSession())
	router.POST("/webhook", controllers.StripeWebhookListener())

	router.POST("/create-payment-intent", controllers.CreatePaymentIntent())
	router.POST("/status-payment-intent", controllers.StatusPaymentIntent())
	router.DELETE("/cancel-payment-intent", controllers.CancelPaymentIntent())
}

func AdminSecurePath(router *gin.Engine) {
	// router.Use(middleware.AdminAuthenticationMiddleware())

	router.GET("/users", controllers.GetUsers())

	// router.POST("/products", controllers.ProductCreate())
	// router.PATCH("/products", controllers.ProductUpdate())
	// router.DELETE("/products", controllers.ProductDelete())
}
