package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "users")

var validate = validator.New()

func Welcome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello programmer..."})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if validationErr := validate.Struct(user); validationErr != nil {
			c.JSON(400, gin.H{"message": validationErr.Error()})
			return
		}

		if count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email}); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		} else if count > 0 {
			c.JSON(400, gin.H{"message": "User allready registerd with given email"})
			return
		}

		if count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone}); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		} else if count > 0 {
			c.JSON(400, gin.H{"message": "User allready registerd with given number"})
			return
		}

		user.ID = primitive.NewObjectID()
		userId := user.ID.Hex()

		user.User_Id = &userId

		password, _ := utils.HashPassword(*user.Password)
		user.Password = &password

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Last_login, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		false := false
		user.IsActive = &false
		user.IsAdmin = &false

		token, refershToken, err := utils.GenerateJWTToken(userId, *user.Email, *user.FirstName, *user.LastName, *user.Phone, *user.IsAdmin, *user.IsActive)

		if err != nil {
			c.JSON(400, gin.H{"message": "Internal Error during GenerateJWTToken"})
			return
		}

		user.Refersh_token = &refershToken
		user.Token = &token

		result, err := userCollection.InsertOne(ctx, user)

		if err != nil {
			fmt.Println("User data not created")
			c.JSON(400, gin.H{"message": "Internal Error"})
			return
		}
		c.JSON(200, gin.H{"message": result})
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "Internal Error"})
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": "User not registed with email"})
			return
		}

		if isPasswordCorrect, _ := utils.VerifyPassword(*user.Password, *foundUser.Password); !isPasswordCorrect {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		userId := foundUser.ID.Hex()

		token, refershToken, _ := utils.GenerateJWTToken(userId, *foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.Phone, *foundUser.IsAdmin, *foundUser.IsActive)

		if err := userCollection.FindOne(ctx, bson.M{"_id": foundUser.ID}).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}
		utils.UpdateLastLogin(userId)

		foundUser.Token = &token
		foundUser.Refersh_token = &refershToken

		c.JSON(200, foundUser)
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if err := utils.CheckUserIsAdmin(c); err != nil {
		// 	c.JSON(401, gin.H{"message": "Unauthorized"})
		// 	return
		// }
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// defer cancel()

		c.JSON(200, gin.H{"message": "Under developement..."})
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		var user models.User

		if err := utils.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(404, gin.H{"message": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}
		c.JSON(200, user)
	}
}
