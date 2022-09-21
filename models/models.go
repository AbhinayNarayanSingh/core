package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName     *string            `json:"firstname" validate:"required"`
	LastName      *string            `json:"lastname" validate:"required"`
	Password      *string            `json:"password" validate:"required"`
	Email         *string            `json:"email" validate:"required"`
	Phone         *string            `json:"phone" validate:"required"`
	IsActive      *bool              `json:"is_active"`
	IsAdmin       *bool              `json:"is_admin"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Last_login    time.Time          `json:"last_login"`
	Token         *string            `json:"token"`
	Refersh_token *string            `json:"refersh_token"`
	User_Id       *string            `json:"user_id"`
}

type OTP struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	OTP_Id  *string            `json:"otp_id"`
	User_Id *string            `json:"user_id"`
	Email   *string            `json:"email"`
	Phone   *string            `json:"phone"`
	OTP     *string            `json:"otp"`
}

type PasswordUpdate struct {
	Email       *string `json:"email"`
	OldPassword *string `json:"old_password" validate:"required"`
	NewPassword *string `json:"new_password" validate:"required"`
	OTP         *string `json:"otp"`
	User_Id     *string `json:"user_id"`
}
