package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Products Models
type Category struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    *string            `bson:"name,omitempty"`
	Icon    *string            `bson:"icon,omitempty"`
	HSN     *string            `bson:"hsn,omitempty"`
	Taxable *string            `bson:"taxable,omitempty"`
	IGST    *string            `bson:"igst,omitempty"`
	CGST    *string            `bson:"cgst,omitempty"`
	SGST    *string            `bson:"sgst,omitempty"`
	Cess    *string            `bson:"cess,omitempty"`
}

type ProductPayload struct {
	ProductDetail
	Varients            *map[string][]Varients `bson:"varients,omitempty"`
	Product_Images      *[]ProductImage        `bson:"product_images,omitempty"`
	Product_Accessories *[]string              `bson:"product_accessories,omitempty"`
	Product_ID          *string                `bson:"product_id,omitempty"`
	Variant_Color       *string                `bson:"variant_color,omitempty"`
	Variant_Size        *string                `bson:"size_size,omitempty"`
	Operation           int                    `json:"operation,omitempty"`
}

type ProductImage struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Product_ID primitive.ObjectID `bson:"product_id,omitempty"`
	Variant    *string            `bson:"variant,omitempty"`
	Images     *[]string          `bson:"images,omitempty"`
}

type ProductVarients struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty"`
	Product_ID primitive.ObjectID     `bson:"product_id,omitempty"`
	Varients   *map[string][]Varients `bson:"varients,omitempty"`
}
type Varients struct {
	Name              *string `bson:"name,omitempty"`
	Description       *string `bson:"description,omitempty"`
	AdditionalDetails *string `bson:"additionalDetails,omitempty"`
	AddonPrice        *int    `bson:"addonPrice,omitempty"`
}

type ProductDetail struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Category       *string            `bson:"category,omitempty"`
	Name           *string            `bson:"name,omitempty"`
	Specifications *[]Varients        `bson:"specifications,omitempty"`
	MRP            *int               `bson:"mrp,omitempty"`
	SellingPrice   *int               `bson:"selling_price,omitempty"`
	Discount       *int               `bson:"discount,omitempty"`
	IsHero         *bool              `bson:"is_hero,omitempty"`
	HeroImage      *bool              `bson:"hero_image,omitempty"`
}
