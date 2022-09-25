package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	FirstName     *string            `bson:"firstname,omitempty" validate:"required"`
	LastName      *string            `bson:"lastname,omitempty" validate:"required"`
	Password      *string            `bson:"password,omitempty" validate:"required"`
	Email         *string            `bson:"email,omitempty" validate:"required"`
	Phone         *string            `bson:"phone,omitempty" validate:"required"`
	IsActive      *bool              `bson:"is_active,omitempty"`
	IsAdmin       *bool              `bson:"is_admin,omitempty"`
	Created_at    time.Time          `bson:"created_at,omitempty"`
	Updated_at    time.Time          `bson:"updated_at,omitempty"`
	Last_login    time.Time          `bson:"last_login,omitempty"`
	Token         *string            `bson:"token,omitempty"`
	Refersh_token *string            `bson:"refersh_token,omitempty"`
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

// Products Models

type Specifications struct {
	Key   *string `bson:"key"`
	Value *string `bson:"value"`
}
type ProductImage struct {
	Product_ID *string   `bson:"product_id,omitempty"`
	Color      *string   `bson:"color,omitempty"`
	Images     *[]string `bson:"product_images,omitempty"`
}

type ProductPrice struct {
	Product_ID    *string `bson:"product_id"`
	Color         *string `bson:"color"`
	MRP           *int    `bson:"mrp"`
	Selling_Price *int    `bson:"selling_price"`
	Discount      *int    `bson:"discount"`
}

type Color struct {
	Color      *string `bson:"color"`
	Color_Code *string `bson:"color_code"`
}

type Product struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Product_Id          *string            `bson:"user_id,omitempty"`
	Category            *string            `bson:"category,omitempty"`
	Name                *string            `bson:"name,omitempty"`
	Color               *[]Color           `bson:"color,omitempty"`
	Size                *[]string          `bson:"size,omitempty"`
	Specifications      *[]Specifications  `bson:"specifications,omitempty"`
	Product_Images      ProductImage       `bson:"product_images,omitempty"`
	Product_Accessories *[]string          `bson:"product_accessories,omitempty"`
	Product_Price       ProductPrice       `bson:"product_price,omitempty"`
}
