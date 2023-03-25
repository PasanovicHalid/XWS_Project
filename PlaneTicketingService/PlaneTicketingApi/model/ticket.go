package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Price float64            `bson:"price" json:"price"`
}
