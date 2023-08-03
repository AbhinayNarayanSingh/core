package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PayloadServices struct {
	ID       string  `json:"_id" bson:"_id"`
	Duration int     `json:"duration" bson:"duration"`
	Name     *string `json:"name,omitempty" bson:"name,omitempty"`
}

type StripePayload struct {
	PaymentIntent_ID string `json:"paymentIntent_id,omitempty" bson:"paymentIntent_id,omitempty"`
}

type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	User_ID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Name        *string            `bson:"name,omitempty" json:"name,omitempty"`
	Address     *string            `bson:"address,omitempty" json:"address,omitempty"`
	City        *string            `bson:"city,omitempty" json:"city,omitempty"`
	State       *string            `bson:"state,omitempty" json:"state,omitempty"`
	Country     *string            `bson:"country,omitempty" json:"country,omitempty"`
	CountryCode *string            `bson:"countryCode,omitempty" json:"countryCode,omitempty"`
	PostalCode  *string            `bson:"postalCode,omitempty" json:"postalCode,omitempty"`
	Email       *string            `bson:"email,omitempty" json:"email,omitempty"`
	Phone       *string            `bson:"phone,omitempty" json:"phone,omitempty"`
	Created_at  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Updated_at  time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Stripe struct {
	Services         []PayloadServices  `json:"services,omitempty" bson:"services,omitempty"`
	BillingAddress   Address            `json:"billingAddress,omitempty" bson:"billingAddress,omitempty"`
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Against          primitive.ObjectID `json:"against,omitempty" bson:"against,omitempty"`
	User_ID          primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Listing_ID       *string            `json:"listing_id,omitempty" bson:"listing_id,omitempty"`
	Seller_ID        *string            `json:"seller_id,omitempty" bson:"seller_id,omitempty"`
	Type             *string            `json:"type" bson:"type"`
	Narration        *string            `json:"narration,omitempty" bson:"narration,omitempty"`
	ClientSecret     string             `json:"clientSecret,omitempty" bson:"clientSecret,omitempty"`
	PaymentIntent_ID string             `json:"paymentIntent_id,omitempty" bson:"paymentIntent_id,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
}
