package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName       *string            `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required"`
	LastName        *string            `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required"`
	Password        *string            `json:"password,omitempty" bson:"password,omitempty" validate:"required"`
	Email           *string            `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
	Phone           *string            `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	IsActive        *bool              `json:"is_active,omitempty" bson:"is_active,omitempty"`
	IsAdmin         *bool              `json:"is_admin,omitempty" bson:"is_admin,omitempty"`
	IsPhoneVerified *bool              `json:"is_phone_verified,omitempty" bson:"is_phone_verified,omitempty"`
	IsEmailVerified *bool              `json:"is_email_verified,omitempty" bson:"is_email_verified,omitempty"`
	Created_at      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Updated_at      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Last_login      time.Time          `json:"last_login,omitempty" bson:"last_login,omitempty"`
	Token           *string            `json:"token,omitempty" bson:"token,omitempty"`
	Telegram_ChatID *string            `json:"telegram_chat_id,omitempty" bson:"telegram_chat_id,omitempty"`
	Operation       int                `json:"operation,omitempty" bson:"operation,omitempty"`
}

type OTP struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User_Id     primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OTP_Id      *string            `json:"otp_id,omitempty" bson:"otp_id,omitempty"`
	Operation   int                `json:"operation,omitempty" bson:"operation,omitempty"`
	Email       *string            `json:"email,omitempty" bson:"email,omitempty"`
	Phone       *string            `json:"phone,omitempty" bson:"phone,omitempty"`
	OTP         *string            `json:"otp,omitempty" bson:"otp,omitempty"`
	OldPassword *string            `json:"old_password,omitempty" bson:"old_password,omitempty"`
	NewPassword *string            `json:"new_password,omitempty" bson:"new_password,omitempty"`
}
