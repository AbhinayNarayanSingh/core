package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/locals"
	"github.com/AbhinayNarayanSingh/core/models"
	"github.com/AbhinayNarayanSingh/core/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "users")
var otpCollection *mongo.Collection = config.OpenCollection(config.Client, "otp")
var resetPasswordCollection *mongo.Collection = config.OpenCollection(config.Client, "reset_password")

var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err.Error()})
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
			c.JSON(400, gin.H{"message": locals.EmailAssociateWithAccount})
			return
		}

		if count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone}); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		} else if count > 0 {
			c.JSON(400, gin.H{"message": locals.EmailNotRegistered})
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

		result, err := userCollection.InsertOne(ctx, user)

		if err != nil {
			fmt.Println("User data not created")
			c.JSON(400, gin.H{"message": locals.InternalServerError})
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
			c.JSON(400, gin.H{"message": locals.InternalServerError})
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": locals.EmailNotRegistered})
			return
		}

		if isPasswordCorrect, _ := utils.VerifyPassword(*user.Password, *foundUser.Password); !isPasswordCorrect {
			c.JSON(401, gin.H{"message": locals.InvalidPassword})
			return
		}

		if !*foundUser.IsActive {
			c.JSON(401, gin.H{"message": locals.AccountNotActivated})
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

func OTPGenerator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var otp models.OTP
		var user models.User
		var updateObject primitive.D

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&otp); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"phone": otp.Phone}).Decode(&user); err != nil {
			c.JSON(400, gin.H{"message": locals.PhoneNotRegistered})
			return
		}

		// here we generate 6 digit otp
		generatedOTP, error := utils.OTPGenerator(6)
		if error != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": error})
			return
		}

		hashOTP, _ := utils.HashPassword(generatedOTP)
		otp.OTP = &hashOTP
		updateObject = append(updateObject, bson.E{Key: "otp", Value: &hashOTP})

		if count, err := otpCollection.CountDocuments(ctx, bson.M{"phone": otp.Phone}); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		} else if count > 0 {
			upsert := true
			filter := bson.M{"phone": otp.Phone}
			update := bson.D{
				{Key: "$set", Value: updateObject},
			}
			opts := options.UpdateOptions{
				Upsert: &upsert,
			}

			if _, err := otpCollection.UpdateOne(ctx, filter, update, &opts); err != nil {
				c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
				return
			}

			c.JSON(200, gin.H{"message": locals.OTPSend, "otp": generatedOTP})
			return
		}

		otp.User_Id = user.User_Id

		otp.ID = primitive.NewObjectID()
		otpHexID := otp.ID.Hex()
		otp.OTP_Id = &otpHexID

		_, err := otpCollection.InsertOne(ctx, otp)

		if err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError})
			return
		}
		c.JSON(200, gin.H{"message": locals.OTPSend, "otp": generatedOTP})
	}
}

func OTPVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var otp models.OTP
		var otpObj models.OTP

		var user models.User
		var updateObject primitive.D

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// post body render
		if err := c.BindJSON(&otp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// finding user object in user model
		if err := userCollection.FindOne(ctx, bson.M{"phone": otp.Phone}).Decode(&user); err != nil {
			c.JSON(400, gin.H{"message": locals.PhoneNotRegistered})
			return
		}

		// searching user phone details in otp models
		if err := otpCollection.FindOne(ctx, bson.M{"phone": otp.Phone}).Decode(&otpObj); err != nil {
			c.JSON(400, gin.H{"message": locals.OTPNotGenerated})
			return
		}

		// verifying otp
		if isOTPCorrect, _ := utils.VerifyPassword(*otp.OTP, *otpObj.OTP); !isOTPCorrect {
			c.JSON(401, gin.H{"message": locals.OTPInvalid})
			return
		}

		updateObject = append(updateObject, bson.E{Key: "isactive", Value: true})

		upsert := true
		filter := bson.M{"phone": otp.Phone}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		// here we're changing status of user account
		if _, err := userCollection.UpdateOne(ctx, filter, bson.D{
			{Key: "$set", Value: updateObject},
		}, &opt); err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		// here we're deleting otp instance from otpCollection
		_, err := otpCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		c.JSON(200, gin.H{"message": locals.AccountActivated})

	}
}

func UpdatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foundUser models.User
		var payload models.PasswordUpdate

		var updateObject primitive.D

		uid := c.GetString("_id")

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// converting json payload to user struct
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": locals.InternalServerError})
			return
		}

		// here we look for user instance
		if err := userCollection.FindOne(ctx, bson.M{"user_id": uid}).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": locals.EmailNotRegistered})
			return
		}

		// here we verify password
		if isPasswordCorrect, _ := utils.VerifyPassword(*payload.OldPassword, *foundUser.Password); !isPasswordCorrect {
			c.JSON(401, gin.H{"message": locals.InvalidPassword})
			return
		}

		// here we hash new password
		if pwd, err := utils.HashPassword(*payload.NewPassword); err != nil {
			c.JSON(401, gin.H{"message": locals.InternalServerError})
			return
		} else {
			updateObject = append(updateObject, bson.E{Key: "password", Value: pwd})
		}

		// here we update password
		upsert := true
		filter := bson.M{"user_id": uid}
		update := bson.D{
			{Key: "$set", Value: updateObject},
		}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		if _, err := userCollection.UpdateOne(ctx, filter, update, &opt); err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		// here we update update_at
		utils.UpdateUpdateAt(uid)

		c.JSON(http.StatusOK, gin.H{"message": "password change sucessfull"})
	}
}

func ResetPasswordInitiator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foundUser models.User
		var payload models.PasswordUpdate

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// first bind payload to struct - email
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": locals.InternalServerError})
		}

		// here we search user instance in userCollection
		if err := userCollection.FindOne(ctx, bson.M{"email": payload.Email}).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": locals.EmailNotRegistered})
			return
		}

		// otp generate
		generatedOTP, error := utils.OTPGenerator(6)
		if error != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": error})
			return
		}
		hashOTP, _ := utils.HashPassword(generatedOTP)
		payload.OTP = &hashOTP
		payload.User_Id = foundUser.User_Id

		// add otp details to PasswordResetCollection
		if _, err := resetPasswordCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": locals.InternalServerError, "details": err})
		}

		c.JSON(200, gin.H{"message": locals.OTPSend, "otp": generatedOTP})
	}
}

func ResetPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.PasswordUpdate

		var foundResetPasswordInstance models.PasswordUpdate

		var updateObject primitive.D

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// first bind payload to struct - email, new_password, otp
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": locals.InternalServerError})
		}

		// here we search user instance in resetPasswordCollection
		if err := resetPasswordCollection.FindOne(ctx, bson.M{"email": payload.Email}).Decode(&foundResetPasswordInstance); err != nil {
			c.JSON(400, gin.H{"message": locals.OTPNotGenerated})
			return
		}

		// here we verify otp
		if isOTPCorrect, _ := utils.VerifyPassword(*payload.OTP, *foundResetPasswordInstance.OTP); !isOTPCorrect {
			c.JSON(401, gin.H{"message": locals.OTPInvalid})
			return
		}

		// here we update user new_password
		pwd, _ := utils.HashPassword(*payload.NewPassword)
		updateObject = append(updateObject, bson.E{Key: "password", Value: pwd})

		upsert := true
		filter := bson.M{"user_id": foundResetPasswordInstance.User_Id}
		update := bson.D{
			{Key: "$set", Value: updateObject},
		}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		if _, err := userCollection.UpdateOne(ctx, filter, update, &opt); err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		// here we update update_at
		utils.UpdateUpdateAt(*foundResetPasswordInstance.User_Id)

		// here we're deleting otp instance from resetPasswordCollection
		_, err := resetPasswordCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "password change sucessfull"})
	}
}