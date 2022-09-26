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

type ProductPayload struct {
	ProductDetail
	Product_Images      *[]ProductImage `bson:"product_images,omitempty"`
	Product_Accessories *[]string       `bson:"product_accessories,omitempty"`
	Product_Price       *[]ProductPrice `bson:"product_price,omitempty"`

	Product_ID    *string `bson:"product_id,omitempty"`
	Variant_Color *string `bson:"variant_color,omitempty"`
	Variant_Size  *string `bson:"size_size,omitempty"`
	Operation     int     `json:"operation,omitempty"`
}

type ProductImage struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Product_ID primitive.ObjectID `bson:"product_id,omitempty"`
	Variant    *string            `bson:"variant,omitempty"`
	Images     *[]string          `bson:"images,omitempty"`
}

type ProductPrice struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty"`
	Product_ID               primitive.ObjectID `bson:"product_id,omitempty"`
	Size                     *string            `bson:"size,omitempty"`
	MRP                      *int               `bson:"mrp,omitempty"`
	Selling_Price            *int               `bson:"selling_price,omitempty"`
	Selling_Price_Before_tax *int               `bson:"selling_price_before_tax,omitempty"`
	Tax                      *int               `bson:"tax,omitempty"`
	Discount                 *int               `bson:"discount,omitempty"`
}

type ProductDetail struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Category       *string            `bson:"category,omitempty"`
	Name           *string            `bson:"name,omitempty"`
	Color          *[]Color           `bson:"color,omitempty"`
	Size           *[]string          `bson:"size,omitempty"`
	Specifications *[]Specifications  `bson:"specifications,omitempty"`
	Product_Price  *[]ProductPrice    `bson:"product_price,omitempty"`
	Product_Images *[]ProductImage    `bson:"product_images,omitempty"`
}

type Color struct {
	Color      *string `bson:"color"`
	Color_Code *string `bson:"color_code"`
}

type Specifications struct {
	Key   *string `bson:"key"`
	Value *string `bson:"value"`
}
