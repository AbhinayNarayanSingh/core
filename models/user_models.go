package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          *string            `bson:"name,omitempty"`
	Type          *string            `bson:"type,omitempty"`
	Number        *string            `bson:"number,omitempty"`
	DisplayNumber *string            `bson:"display_number,omitempty"`
	ExpDate       *string            `bson:"exp_date,omitempty"`
	Created_at    time.Time          `bson:"created_at,omitempty"`
}
