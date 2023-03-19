package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	Id                  primitive.ObjectID
	StartDateTimeUTC    time.Time
	EndDateTimeUTC      time.Time
	DepartureLocation   string
	DestinationLocation string
	Price               float64
	AvailableTickets    []Ticket
}
