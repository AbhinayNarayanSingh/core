package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	User_ID     primitive.ObjectID `bson:"user_id,omitempty"`
	FirstName   *string            `bson:"firstname,omitempty"`
	LastName    *string            `bson:"lastname,omitempty"`
	Address     *string            `bson:"address,omitempty"`
	City        *string            `bson:"city,omitempty"`
	State       *string            `bson:"state,omitempty"`
	StateCode   *string            `bson:"stateCode,omitempty"`
	Country     *string            `bson:"country,omitempty"`
	CountryCode *string            `bson:"countryCode,omitempty"`
	PostalCode  *string            `bson:"postalCode,omitempty"`
	Email       *string            `bson:"email,omitempty"`
	Phone       *string            `bson:"phone,omitempty"`
	Created_at  time.Time          `bson:"created_at,omitempty"`
	Updated_at  time.Time          `bson:"updated_at,omitempty"`
}

type Card struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          *string            `bson:"name,omitempty"`
	Type          *string            `bson:"type,omitempty"`
	Number        *string            `bson:"number,omitempty"`
	DisplayNumber *string            `bson:"display_number,omitempty"`
	ExpDate       *string            `bson:"exp_date,omitempty"`
	Created_at    time.Time          `bson:"created_at,omitempty"`
}
