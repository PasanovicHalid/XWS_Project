package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id                  primitive.ObjectID `bson:"_id" json:"id"`
	StartDateTimeUTC    time.Time          `bson:"_id" json:"id"`
	EndDateTimeUTC      time.Time          `bson:"_id" json:"id"`
	DepartureLocation   string             `bson:"_id" json:"id"`
	DestinationLocation string             `bson:"_id" json:"id"`
	Price               float64            `bson:"_id" json:"id"`
	AvailableTickets    []Ticket           `bson:"_id" json:"id"`
}
