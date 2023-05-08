package domain

import (
	"time"
)

type AccommodationOffer struct {
	Id               string        `bson:"_id,omitempty"`
	AccommodationId  Accommodation `bson:"accommodation,omitempty"`
	StartDateTimeUTC time.Time     `bson:"startDateTimeUTC" json:"startDateTimeUTC"`
	EndDateTimeUTC   time.Time     `bson:"endDateTimeUTC" json:"endDateTimeUTC"`
	Price            int           `bson:"price,omitempty"`
}
