package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	User_ID           primitive.ObjectID `bson:"user_id,omitempty"`
	Address_ID        primitive.ObjectID `bson:"address_id,omitempty"`
	Transaction_ID    primitive.ObjectID `bson:"transaction_id,omitempty"`
	Tracking_ID       primitive.ObjectID `bson:"tracking_id,omitempty"`
	OrderNumber       *string            `bson:"orderNumber,omitempty"`
	Total             *int               `bson:"total,omitempty"`
	Shipping          *int               `bson:"shipping,omitempty"`
	Tax               *int               `bson:"tax,omitempty"`
	Quantity          *int               `bson:"quantity,omitempty"`
	IsOrderConfirmed  *bool              `bson:"isOrderConfirmed,omitempty"`
	IsDelivered       *bool              `bson:"isDelivered,omitempty"`
	Created_at        time.Time          `bson:"created_at,omitempty"`
	OrderPlaced_at    time.Time          `bson:"order_placed_at,omitempty"`
	Shipped_at        time.Time          `bson:"shipped_at,omitempty"`
	OutForDelivery_at time.Time          `bson:"out_for_delivery_at,omitempty"`
	Delivery_at       time.Time          `bson:"delivery_at,omitempty"`
}

type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Status         *string            `bson:"status,omitempty"`
	Amount         *string            `bson:"amount,omitempty"`
	Transaction_ID *string            `bson:"transaction_id,omitempty"`
	Currency       *string            `bson:"currency,omitempty"`
	Created_at     time.Time          `bson:"created_at,omitempty"`
}

type Tracking struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	Carrier                 *string            `bson:"carrier,omitempty"`
	Method                  *string            `bson:"method,omitempty"`
	Shipping                *int               `bson:"shipping,omitempty"`
	Weight                  *int               `bson:"weight,omitempty"`
	Description             *string            `bson:"description,omitempty"`
	IsInternationalDelivery *bool              `bson:"international_delivery,omitempty"`
	Expected_Delivery_at    time.Time          `bson:"expected_delivery,omitempty"`
	Created_at              time.Time          `bson:"created_at,omitempty"`
	Updated_at              time.Time          `bson:"updated_at,omitempty"`
}
