package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id                  primitive.ObjectID `bson:"_id" json:"id"`
	StartDateTimeUTC    time.Time          `bson:"startDateTimeUTC" json:"startDateTimeUTC"`
	EndDateTimeUTC      time.Time          `bson:"endDateTimeUTC" json:"endDateTimeUTC"`
	DepartureLocation   string             `bson:"departureLocation " json:"departureLocation"`
	DestinationLocation string             `bson:"destinationLocation" json:"destinationLocation"`
	Price               float64            `bson:"price" json:"price"`
	AvailableTickets    []Ticket           `bson:"availableTickets" json:"availableTickets"`
}
