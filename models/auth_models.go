package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	FirstName       *string            `bson:"firstname,omitempty" validate:"required"`
	LastName        *string            `bson:"lastname,omitempty" validate:"required"`
	Password        *string            `bson:"password,omitempty" validate:"required"`
	Email           *string            `bson:"email,omitempty" validate:"required"`
	Phone           *string            `bson:"phone,omitempty" validate:"required"`
	IsActive        *bool              `bson:"is_active,omitempty"`
	IsAdmin         *bool              `bson:"is_admin,omitempty"`
	Created_at      time.Time          `bson:"created_at,omitempty"`
	Updated_at      time.Time          `bson:"updated_at,omitempty"`
	Last_login      time.Time          `bson:"last_login,omitempty"`
	Token           *string            `bson:"token,omitempty"`
	Telegram_ChatID *string            `bson:"telegram_chat_id,omitempty"`
	Operation       int                `bson:"operation,omitempty"`
}

type OTP struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	User_Id primitive.ObjectID `bson:"user_id,omitempty"`
	Email   *string            `bson:"email,omitempty"`
	Phone   *string            `bson:"phone,omitempty"`
	OTP     *string            `bson:"otp,omitempty"`
}

type PasswordUpdate struct {
	User_Id     primitive.ObjectID `bson:"user_id,omitempty"`
	Email       *string            `bson:"email,omitempty"`
	OldPassword *string            `bson:"old_password,omitempty" validate:"required"`
	NewPassword *string            `bson:"new_password,omitempty" validate:"required"`
	OTP         *string            `bson:"otp,omitempty"`
}
