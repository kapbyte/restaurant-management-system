package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID primitive.ObjectID `bson:"_id"`
}
