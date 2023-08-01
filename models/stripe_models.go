package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PayloadServices struct {
	Service       *string  `json:"service" bson:"service"`
	Vadility      *int     `json:"vadility" bson:"vadility"`
	UnitOfMeasure *string  `json:"unitOfMeasure" bson:"unitOfMeasure"`
	BasePrice     *float32 `json:"basePrice" bson:"basePrice"`
}

type Stripe struct {
	Services         []PayloadServices  `json:"services" bson:"services"`
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Against          primitive.ObjectID `json:"against,omitempty" bson:"against,omitempty"`
	User_ID          primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Listing_ID       *string            `json:"listing_id,omitempty" bson:"listing_id,omitempty"`
	Seller_ID        *string            `json:"seller_id,omitempty" bson:"seller_id,omitempty"`
	Type             *string            `json:"type" bson:"type"`
	Narration        *string            `json:"narration,omitempty" bson:"narration,omitempty"`
	Amount           *float32           `json:"amount" bson:"amount"`
	ClientSecret     string             `json:"clientSecret,omitempty" bson:"clientSecret,omitempty"`
	PaymentIntent_ID string             `json:"paymentIntent_id,omitempty" bson:"paymentIntent_id,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
}
