package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	Id    primitive.ObjectID
	Price float64
}
