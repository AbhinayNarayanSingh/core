package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	User_ID    primitive.ObjectID `bson:"user_id,omitempty"`
	FirstName  *string            `bson:"firstname,omitempty" validate:"required"`
	LastName   *string            `bson:"lastname,omitempty" validate:"required"`
	Address    *string            `bson:"address,omitempty" validate:"required"`
	City       *string            `bson:"city,omitempty" validate:"required"`
	State      *string            `bson:"state,omitempty" validate:"required"`
	StateCode  *string            `bson:"stateCode,omitempty" validate:"required"`
	Country    *string            `bson:"country,omitempty" validate:"required"`
	PostalCode *string            `bson:"postalCode,omitempty" validate:"required"`
	Email      *string            `bson:"email,omitempty" validate:"required"`
	Phone      *string            `bson:"phone,omitempty" validate:"required"`
	Created_at time.Time          `bson:"created_at,omitempty"`
	Updated_at time.Time          `bson:"updated_at,omitempty"`
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
