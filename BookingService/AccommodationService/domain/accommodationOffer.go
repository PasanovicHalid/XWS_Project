package domain

import (
	"time"
)

type AccommodationOffer struct {
	Id                        string    `bson:"_id,omitempty"`
	Deleted                   bool      `bson:"deleted"`
	SagaTimestamp             int64     `bson:"saga_timestamp"`
	AccommodationId           string    `bson:"accommodationId,omitempty"`
	AvailableStartDateTimeUTC time.Time `bson:"AvailabletartDateTimeUTC" json:"startDateTimeUTC"`
	AvailableEndDateTimeUTC   time.Time `bson:"AvailablendDateTimeUTC" json:"endDateTimeUTC"`
	Price                     int       `bson:"price,omitempty"`
	PerGuest                  bool      `bson:"perGuest,omitempty"`
	AutomaticAcceptation      bool      `bson:"automatic"`
}
