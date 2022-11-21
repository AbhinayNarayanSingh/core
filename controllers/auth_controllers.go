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
			c.JSON(400, gin.H{"message": locals.PhoneAssociateWithAccount})
			return
		}

		user.ID = primitive.NewObjectID()

		password, _ := utils.HashPassword(*user.Password)
		user.Password = &password

		user.Created_at = utils.TimeStampFn()

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
		var payload models.User
		var foundUser models.User

		payload.Operation = 1

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.BadRequest})
			return
		}

		filter := bson.M{}
		msg := locals.EmailNotRegistered
		switch payload.Operation {
		case 1:
			filter = bson.M{"email": payload.Email}
		case 2:
			filter = bson.M{"phone": payload.Phone}
			msg = locals.PhoneNotRegistered
		}

		if err := userCollection.FindOne(ctx, filter).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": msg})
			return
		}

		if isPasswordCorrect, _ := utils.VerifyPassword(*payload.Password, *foundUser.Password); !isPasswordCorrect {
			c.JSON(401, gin.H{"message": locals.InvalidPassword})
			return
		}

		if !*foundUser.IsActive {
			c.JSON(401, gin.H{"message": locals.AccountNotActivated})
			return
		}

		userId := foundUser.ID.Hex()

		token, _ := utils.GenerateJWTToken(userId, *foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.Phone, *foundUser.IsAdmin, *foundUser.IsActive)

		utils.UpdateTimeStampFn(userCollection, &foundUser.ID, "last_login")

		foundUser.Token = &token

		c.JSON(200, foundUser)
		// text := "Your OTP is " + "157379" + " for login on iCorn, never share your code with anyone."
		text := "Hello " + *foundUser.FirstName + ", We detected a login to your account"
		utils.SendTelegramMessage(*foundUser.Telegram_ChatID, text)
	}
}

// 1.email verification		2.phone number verification		3.signup account activation with phone		4.mobile sign in
// 5.email password reset		6. phone password reset

func OTPVerificationInitiator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.OTP

		var user models.User
		var updateObject primitive.D

		payload.Operation = 1

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		filter := bson.M{}
		msg := locals.PhoneNotRegistered
		sucessMsg := locals.OTPSendOnPhone

		switch payload.Operation {
		case 1, 5:
			filter = bson.M{"email": payload.Email}
			msg = locals.EmailNotRegistered
		case 2, 3, 4, 6:
			filter = bson.M{"phone": payload.Phone}
		}

		switch payload.Operation {
		case 1:
			sucessMsg = locals.OTPSendOnEmail
		case 5:
			sucessMsg = locals.OTPSendOnEmailForReset
		case 6:
			sucessMsg = locals.OTPSendOnPhoneForReset
		}

		if err := userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
			c.JSON(400, gin.H{"message": msg})
			return
		}

		// here we generate 6 digit otp
		hashOTP, generatedOTP, error := utils.OTPGenerator(6)
		if error != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": error})
			return
		}

		payload.OTP = &hashOTP

		otp_id := primitive.NewObjectID()
		payload.ID = otp_id
		otp_id_string := otp_id.Hex()

		updateObject = append(updateObject, bson.E{Key: "otp", Value: &hashOTP})
		updateObject = append(updateObject, bson.E{Key: "operation", Value: payload.Operation})

		if count, err := otpCollection.CountDocuments(ctx, filter); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		} else if count > 0 {
			upsert := true
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

			c.JSON(200, gin.H{"message": sucessMsg, "otp": generatedOTP, "otp_id": otp_id_string})
			text := "Hello " + *user.FirstName + ", your security code is " + generatedOTP + ". never share your OTP with anyone else!"
			go utils.SendTelegramMessage(*user.Telegram_ChatID, text)
			return
		}

		payload.User_Id = user.ID

		if _, err := otpCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError})
			return
		}
		c.JSON(200, gin.H{"message": sucessMsg, "otp": generatedOTP, "otp_id": otp_id_string})
	}
}

func OTPVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.OTP
		payload.Operation = 3

		var user models.User

		var otpObj models.OTP
		var updateObject primitive.D

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		lookup := bson.M{"_id": payload.OTP_Id}

		var DeleteOTPInstanceFn = func() {
			_, err := otpCollection.DeleteOne(ctx, lookup)
			if err != nil {
				c.JSON(400, gin.H{"message": err})
				return
			}
		}

		if err := otpCollection.FindOne(ctx, lookup).Decode(&otpObj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.OTPNotGenerated})
			return
		}

		if isOTPCorrect, _ := utils.VerifyPassword(*payload.OTP, *otpObj.OTP); !isOTPCorrect {
			c.JSON(401, gin.H{"message": locals.OTPInvalid})
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"_id": otpObj.User_Id}).Decode(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.UserNotFound})
			return
		}

		var collection *mongo.Collection

		switch payload.Operation {
		case 3:
			collection = userCollection
			upsert := true
			filter := bson.M{"_id": otpObj.User_Id}
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			updateObject = append(updateObject, bson.E{Key: "is_active", Value: true})
			update := bson.D{
				{Key: "$set", Value: updateObject},
			}

			if _, err := collection.UpdateOne(ctx, filter, update, &opt); err != nil {
				c.JSON(400, gin.H{"message": err})
				return
			}

			c.JSON(200, gin.H{"message": locals.AccountActivated})
			go DeleteOTPInstanceFn()
			return
		}
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
		utils.UpdateTimeStampFn(userCollection, &foundUser.ID, "updated_at")

		c.JSON(http.StatusOK, gin.H{"message": "password change sucessfull"})
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
		fmt.Println(1)
		// here we verify otp
		if isOTPCorrect, _ := utils.VerifyPassword(*payload.OTP, *foundResetPasswordInstance.OTP); !isOTPCorrect {
			c.JSON(401, gin.H{"message": locals.OTPInvalid})
			return
		}

		fmt.Println(2)
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

		fmt.Println(3)
		if _, err := userCollection.UpdateOne(ctx, filter, update, &opt); err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		fmt.Println(4)
		// here we update update_at
		utils.UpdateTimeStampFn(userCollection, &foundResetPasswordInstance.User_Id, "updated_at")

		fmt.Println(5)
		// here we're deleting otp instance from resetPasswordCollection
		_, err := resetPasswordCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}

		fmt.Println(6)
		c.JSON(http.StatusOK, gin.H{"message": "password change sucessfull"})
	}
}
