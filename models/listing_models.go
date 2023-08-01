package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        *string            `bson:"name,omitempty" json:"name,omitempty"`
	Icon        *string            `bson:"icon,omitempty" json:"icon,omitempty"`
	SubCategory *[]SubCategory     `bson:"sub_category" json:"sub_category"`
}

type SubCategory struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string            `bson:"name,omitempty" json:"name,omitempty"`
	Icon *string            `bson:"icon,omitempty" json:"icon,omitempty"`
}

type Services struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type          *string            `json:"type,omitempty" bson:"type,omitempty"`
	Name          *string            `json:"name,omitempty" bson:"name,omitempty"`
	Label         *string            `json:"label,omitempty" bson:"label,omitempty"`
	BasePrice     float32            `json:"basePrice" bson:"basePrice"`
	Vadility      []int              `json:"vadility" bson:"vadility"`
	UnitOfMeasure *string            `json:"unitOfMeasure" bson:"unitOfMeasure"`
}

type Listing struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	Category_Id             primitive.ObjectID `bson:"category_id,omitempty"`
	Seller_Id               primitive.ObjectID `bson:"seller_id,omitempty"`
	UID                     int                `bson:"uid,omitempty"`
	Slug                    string             `bson:"slug,omitempty"`
	Category                string             `bson:"category,omitempty"`
	Title                   string             `bson:"title"`
	Description             string             `bson:"description"`
	Listing_for             string             `bson:"listing_for"`
	Amount                  float32            `bson:"amount"`
	Currency                string             `bson:"currency"`
	Images                  *[]Image           `bson:"images"`
	IsActiveAd              bool               `bson:"isActiveAd"`
	IsFeaturedAd            bool               `bson:"isFeaturedAd"`
	IsHighlightAd           bool               `bson:"isHighlightAd"`
	IsBumpAd                bool               `bson:"isBumpAd"`
	IsWebsiteLinkedAd       bool               `bson:"isWebsiteLinkedAd"`
	Tags                    *[]string          `bson:"tags"`
	WebsiteURL              string             `bson:"websiteURL"`
	YoutubeVideoURL         string             `bson:"youtubeVideoURL"`
	Seller                  string             `bson:"seller"`
	Formatted_address       string             `bson:"formatted_address"`
	Short_formatted_address string             `bson:"short_formatted_address"`
	Place_id                string             `bson:"place_id"`
	Lat                     float64            `bson:"lat"`
	Lng                     float64            `bson:"lng"`
	Country_code            string             `bson:"country_code"`
	Phone                   string             `bson:"phone"`
	Email                   string             `bson:"email"`
	Posted_on               time.Time          `bson:"posted_on"`
	Updated_on              time.Time          `bson:"updated_on"`
	AdExpiry                time.Time          `bson:"ad_expiry"`
	FeaturedExp             time.Time          `bson:"featured_exp"`
	HighlightExp            time.Time          `bson:"highlight_exp"`
	BumpExp                 time.Time          `bson:"bump_exp"`
}

type Image struct {
	URL       string `bson:"url"`
	Public_Id string `bson:"public_id"`
}
