package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Products Models
type Category struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

type ProductPayload struct {
	ProductDetail
	Varients       *Varients       `bson:"varients,omitempty"`
	Product_Images *[]ProductImage `bson:"product_images,omitempty"`

	Product_Accessories *[]string `bson:"product_accessories,omitempty"`
	Product_ID          *string   `bson:"product_id,omitempty"`
	Variant_Color       *string   `bson:"variant_color,omitempty"`
	Variant_Size        *string   `bson:"size_size,omitempty"`
	Operation           int       `json:"operation,omitempty"`
}

type ProductImage struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Product_ID primitive.ObjectID `bson:"product_id,omitempty"`
	Variant    *string            `bson:"variant,omitempty"`
	Images     *[]string          `bson:"images,omitempty"`
}

type Varients struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Product_ID   primitive.ObjectID `bson:"product_id,omitempty"`
	Storage      *[]ProductVarients `bson:"storage,omitempty"`
	Finish       *[]ProductVarients `bson:"finish,omitempty"`
	Model        *[]ProductVarients `bson:"model,omitempty"`
	Memory       *[]ProductVarients `bson:"memory,omitempty"`
	Processor    *[]ProductVarients `bson:"processor,omitempty"`
	Connectivity *[]ProductVarients `bson:"connectivity,omitempty"`
}
type ProductVarients struct {
	VarientName              *string `bson:"varientName,omitempty"`
	VarientDiscription       *string `bson:"varientDiscription,omitempty"`
	VarientAdditionalDetails *string `bson:"varientAdditionalDetails,omitempty"`
	VarientAddonPrice        *int    `bson:"varientAddonPrice,omitempty"`
}

type ProductDetail struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Category       *string            `bson:"category,omitempty"`
	Name           *string            `bson:"name,omitempty"`
	Specifications *[]Specifications  `bson:"specifications,omitempty"`
	MRP            *int               `bson:"mrp,omitempty"`
	Selling_Price  *int               `bson:"selling_price,omitempty"`
	Discount       *int               `bson:"discount,omitempty"`
}

type Color struct {
	Color      *string `bson:"color"`
	Color_Code *string `bson:"color_code"`
}

type Specifications struct {
	Key   *string `bson:"key"`
	Value *string `bson:"value"`
}
