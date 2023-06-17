package models

import (
	"context"
	"fmt"
	"time"

	"github.com/AbhinayNarayanSingh/core/config"
	"github.com/AbhinayNarayanSingh/core/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "users")
var otpCollection *mongo.Collection = config.OpenCollection(config.Client, "otp")

type User struct {
	ID                    primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName             *string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required"`
	LastName              *string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required"`
	Password              *string             `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Email                 *string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	Phone                 *string             `json:"phone,omitempty" bson:"phone"`
	IsActive              *bool               `json:"is_active,omitempty" bson:"is_active,omitempty"`
	IsAdmin               *bool               `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
	IsPhoneVerified       *bool               `json:"is_phone_verified,omitempty" bson:"is_phone_verified,omitempty"`
	IsEmailVerified       *bool               `json:"is_email_verified,omitempty" bson:"is_email_verified,omitempty"`
	Created_at            time.Time           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Updated_at            time.Time           `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Last_login            time.Time           `json:"last_login,omitempty" bson:"last_login,omitempty"`
	Token                 *string             `json:"token,omitempty" bson:"token,omitempty"`
	ReferenceToken        *string             `json:"reference_token,omitempty" bson:"reference_token,omitempty"`
	Telegram_ChatID       *string             `json:"telegram_chat_id,omitempty" bson:"telegram_chat_id,omitempty"`
	Operation             int                 `json:"operation,omitempty" bson:"operation,omitempty"`
	IsProfessionalAccount *bool               `json:"is_professional_account,omitempty" bson:"is_professional_account,omitempty"`
	ProfessionalDetail    *ProfessionalDetail `json:"professional_detail,omitempty" bson:"professional_detail,omitempty"`
}

type ProfessionalDetail struct {
	Logo        *string `json:"logo,omitempty" bson:"logo,omitempty"`
	Name        *string `json:"name,omitempty" bson:"name,omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	Website     *string `json:"website,omitempty" bson:"website,omitempty"`
	Email       *string `json:"email,omitempty" bson:"email,omitempty"`
	Contact     *string `json:"contact,omitempty" bson:"contact,omitempty"`
}

func (user User) PasswordVerify(userEnteredPassword string) bool {
	return utils.VerifyPassword(userEnteredPassword, *user.Password)
}

func (user User) AccessToken() (string, string) {
	userId := user.ID.Hex()
	token, _ := utils.GenerateJWTToken(userId, *user.Email, *user.FirstName, *user.LastName, *user.Phone, *user.IsAdmin, *user.IsActive, 7)
	referenceToken, _ := utils.GenerateJWTToken(userId, *user.Email, *user.FirstName, *user.LastName, *user.Phone, *user.IsAdmin, *user.IsActive, 180)
	utils.UpdateTimeStampFn(userCollection, &user.ID, "last_login")
	return token, referenceToken
}

type OTP struct { // payload
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_Id     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OTP_Id      *string            `json:"otp_id,omitempty" bson:"otp_id,omitempty"`
	Operation   int                `json:"operation,omitempty" bson:"operation,omitempty"`
	Email       *string            `json:"email,omitempty" bson:"email,omitempty"`
	Phone       *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	OTP         *string            `json:"otp,omitempty" bson:"otp,omitempty"`
	Password    *string            `json:"password,omitempty" bson:"password,omitempty"`
	OldPassword *string            `json:"old_password,omitempty" bson:"old_password,omitempty"`
	NewPassword *string            `json:"new_password,omitempty" bson:"new_password,omitempty"`
}

func (otp OTP) OTPVerify(userEnteredOTP string) bool {
	return utils.VerifyPassword(userEnteredOTP, *otp.OTP)
}

func (otp OTP) RemoveOTP() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lookup := bson.M{"_id": otp.ID}
	if _, err := otpCollection.DeleteMany(ctx, lookup); err != nil {
		fmt.Println(err)
	}
	return
}
