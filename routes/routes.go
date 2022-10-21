package routes

import (
	"github.com/AbhinayNarayanSingh/core/controllers"
	"github.com/AbhinayNarayanSingh/core/middleware"
	"github.com/gin-gonic/gin"
)

func Path(router *gin.Engine) {

	router.GET("/", controllers.Welcome())
	router.POST("/signup", controllers.SignUp())
	router.POST("/signin", controllers.SignIn())

	router.POST("/signup/otp", controllers.OTPGenerator())
	router.POST("/signup/otp/verify", controllers.OTPVerify())

	router.POST("/user/password/reset", controllers.ResetPasswordInitiator())
	router.POST("/user/password/reset/verify", controllers.ResetPassword())

	router.POST("/product", controllers.ProductGet())
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
