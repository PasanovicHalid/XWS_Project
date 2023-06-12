package domain

import "time"

type Rating struct {
	Id              string    `bson:"_id,omitempty"`
	UserId          string    `bson:"userId,omitempty"`
	HostId          string    `bson:"hostId,omitempty"`
	AccommodationId string    `bson:"accommodationId,omitempty"`
	Rating          float64   `bson:"rating,omitempty"`
	TimeIssued      time.Time `bson:"timestamp,omitempty"`
}
