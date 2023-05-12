package domain

import (
	"time"
)

type AccommodationOffer struct {
	Id                        string    `bson:"_id,omitempty"`
	AccommodationId           string    `bson:"accommodationId,omitempty"`
	AvailableStartDateTimeUTC time.Time `bson:"AvailabletartDateTimeUTC" json:"startDateTimeUTC"`
	AvailableEndDateTimeUTC   time.Time `bson:"AvailablendDateTimeUTC" json:"endDateTimeUTC"`
	Price                     int       `bson:"price,omitempty"`
}
