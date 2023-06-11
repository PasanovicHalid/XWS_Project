package domain

import (
	"time"
)

type Reservation struct {
	Id                   string            `bson:"_id,omitempty"`
	Deleted              bool              `bson:"deleted"`
	SagaTimestamp        int64             `bson:"saga_timestamp"`
	AccommodationOfferId string            `bson:"offerId,omitempty"`
	CustomerId           string            `bson:"customerId,omitempty"`
	HostId               string            ` bson:"hostId,omitempty"`
	ReservationStatus    ReservationStatus `bson:"reservationStatus"`
	NumberOfGuests       int               `bson:"numOfGuests,omitempty"`
	StartDateTimeUTC     time.Time         `bson:"startDateTimeUTC" json:"startDateTimeUTC"`
	EndDateTimeUTC       time.Time         `bson:"endDateTimeUTC" json:"endDateTimeUTC"`
}

type ReservationStatus int

const (
	Pending ReservationStatus = iota
	Accepted
	Rejected
	Canceled
)
