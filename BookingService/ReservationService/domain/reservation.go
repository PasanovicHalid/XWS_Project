package domain

import (
	"time"
)

type Reservation struct {
	Id                   string            `bson:"_id,omitempty"`
	AccommodationOfferId string            `bson:"offerId,omitempty"`
	CustomerId           string            `bson:"customerId,omitempty"`
	Status               ReservationStatus `bson:"status,omitempty"`
	NumberOfGuests       int               `bson:"numOfGuests,omitempty"`
	StartDateTimeUTC     time.Time         `bson:"startDateTimeUTC" json:"startDateTimeUTC"`
	EndDateTimeUTC       time.Time         `bson:"endDateTimeUTC" json:"endDateTimeUTC"`
}
type ReservationStatus int

const (
	Pending ReservationStatus = iota
	Accepted
	Rejected
)
