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
var OtpCollection *mongo.Collection = config.OpenCollection(config.Client, "otp")
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
		user.IsEmailVerified = &false
		user.IsPhoneVerified = &false

		result, err := userCollection.InsertOne(ctx, user)

		if err != nil {
			fmt.Println("User data not created")
			c.JSON(400, gin.H{"message": locals.InternalServerError, "err": err})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": result.InsertedID})

		text := "Hello " + *user.FirstName + locals.AccountCreated
		go utils.SendTelegramMessage(*user.Telegram_ChatID, text)
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.OTP

		var foundOTP models.OTP
		var foundUser models.User

		payload.Operation = 1

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"message": locals.BadRequest})
			return
		}

		filter := bson.M{}
		var msg string

		switch payload.Operation {
		case 1:
			filter = bson.M{"email": payload.Email}
			msg = locals.EmailNotRegistered

		case 2:
			filter = bson.M{"phone": payload.Phone}
			msg = locals.PhoneNotRegistered

		case 4:
			otp_id, _ := primitive.ObjectIDFromHex(*payload.OTP_Id)
			lookup := bson.M{"_id": otp_id}
			if err := OtpCollection.FindOne(ctx, lookup).Decode(&foundOTP); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
			filter = bson.M{"_id": foundOTP.User_Id}

		default:
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.BadRequest})
			return
		}

		if err := userCollection.FindOne(ctx, filter).Decode(&foundUser); err != nil {
			c.JSON(400, gin.H{"message": msg})
			return
		}

		switch payload.Operation {
		case 1, 2:
			if isVerified := foundUser.PasswordVerify(*payload.Password); !isVerified {
				c.JSON(401, gin.H{"message": locals.InvalidPassword})
				return
			}
		case 4:
			if isVerified := foundOTP.OTPVerify(*payload.OTP); !isVerified {
				c.JSON(401, gin.H{"message": locals.InvalidPassword})
				return
			}
		}

		if !*foundUser.IsActive {
			c.JSON(401, gin.H{"message": locals.AccountNotActivated})
			return
		}

		token := foundUser.AccessToken()
		foundUser.Token = &token

		stringempty := ""
		foundUser.Password = &stringempty

		c.JSON(200, foundUser)
		text := "Hello " + *foundUser.FirstName + ", We detected a login to your account"
		go utils.SendTelegramMessage(*foundUser.Telegram_ChatID, text)
		if payload.Operation == 4 {
			go foundOTP.RemoveOTP()
		}
	}
}

// 1.email verification		2.phone number verification
// 3.signup account activation with phone
// 5.email password reset		6. phone password reset
// 7. sigin with email password				8.sigin with mobile password		4.mobile sign in with OTP

func OTPVerificationInitiator(action int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.OTP
		var user models.User

		payload.Operation = 3
		if action != 0 {
			payload.Operation = action
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		filter := bson.M{}
		errorMsg := ""
		sucessMsg := locals.OTPSendOnPhone

		// otp send on email or phone
		switch payload.Operation {
		case 1, 5:
			filter = bson.M{"email": payload.Email}
			errorMsg = locals.EmailNotRegistered
		case 2, 3, 4, 6:
			filter = bson.M{"phone": payload.Phone}
			errorMsg = locals.PhoneNotRegistered
		}

		// response message for client
		switch payload.Operation {
		case 1:
			sucessMsg = locals.OTPSendOnEmail
		case 5:
			sucessMsg = locals.OTPSendOnEmailForReset
		case 6:
			sucessMsg = locals.OTPSendOnPhoneForReset
		}

		// looking for user
		if err := userCollection.FindOne(ctx, filter).Decode(&user); err != nil {
			c.JSON(400, gin.H{"message": errorMsg})
			return
		}

		// here we generate 6 digit otp
		hashOTP, generatedOTP, error := utils.OTPGenerator(6)
		if error != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": error})
			return
		}

		payload.OTP = &hashOTP

		otp_uid := primitive.NewObjectID()
		payload.ID = otp_uid

		payload.User_Id = user.ID

		otp_id_string := otp_uid.Hex()

		if _, err := OtpCollection.DeleteMany(ctx, filter); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError, "details": err})
			return
		}

		if _, err := OtpCollection.InsertOne(ctx, payload); err != nil {
			c.JSON(400, gin.H{"message": locals.InternalServerError})
			return
		}

		c.JSON(200, gin.H{"message": sucessMsg, "otp": generatedOTP, "otp_id": otp_id_string})

		if user.Telegram_ChatID != nil {
			text := "Hello " + *user.FirstName + ", your security code is " + generatedOTP + ". never share your OTP with anyone else!"

			go utils.SendTelegramMessage(*user.Telegram_ChatID, text)
		}
		return
	}
}

func OTPVerification(action int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.OTP
		payload.Operation = 3
		if action != 0 {
			payload.Operation = action
		}

		var user models.User

		var otpObj models.OTP
		var updateObject primitive.D

		var collection *mongo.Collection
		var sucessMsg string

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		id, err := primitive.ObjectIDFromHex(*payload.OTP_Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		lookup := bson.M{"_id": id}

		if err := OtpCollection.FindOne(ctx, lookup).Decode(&otpObj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.OTPNotGenerated})
			return
		}

		if isOTPCorrect := otpObj.OTPVerify(*payload.OTP); !isOTPCorrect {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.OTPInvalid})
			return
		}

		if err := userCollection.FindOne(ctx, bson.M{"_id": otpObj.User_Id}).Decode(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": locals.UserNotFound})
			return
		}

		filter := bson.M{"_id": otpObj.User_Id}

		updateObject = append(updateObject, bson.E{Key: "updated_at", Value: utils.TimeStampFn()})

		switch otpObj.Operation {
		case 1:
			updateObject = append(updateObject, bson.E{Key: "is_email_verified", Value: true})
			sucessMsg = "We have sucessfully verified your email address"

		case 2:
			updateObject = append(updateObject, bson.E{Key: "is_phone_verified", Value: true})
			sucessMsg = "We have sucessfully verified your phone number"

		case 3:
			updateObject = append(updateObject, bson.E{Key: "is_active", Value: true})
			updateObject = append(updateObject, bson.E{Key: "is_phone_verified", Value: true})
			sucessMsg = locals.AccountActivated

		case 5, 6:
			pwd, _ := utils.HashPassword(*payload.NewPassword)
			updateObject = append(updateObject, bson.E{Key: "password", Value: pwd})
			sucessMsg = "Password has been sucessfully reset"
		}

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		update := bson.D{
			{Key: "$set", Value: updateObject},
		}
		collection = userCollection

		switch otpObj.Operation {
		case 4:
			c.JSON(http.StatusBadRequest, gin.H{"message": "StatusBadRequest"})
			return

		case 1, 2, 3, 5:
			if _, err := collection.UpdateOne(ctx, filter, update, &opt); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
				return
			}
			c.JSON(200, gin.H{"message": sucessMsg})
		}

		go utils.SendTelegramMessage(*user.Telegram_ChatID, sucessMsg)
		_, err = OtpCollection.DeleteMany(ctx, lookup)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			return
		}
	}
}

func UpdatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foundUser models.User
		var payload models.OTP

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
		if isPasswordCorrect := payload.OTPVerify(*foundUser.Password); !isPasswordCorrect {
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
