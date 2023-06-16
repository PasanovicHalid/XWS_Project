package domain

import "time"

type Rating struct {
	Id              string    `bson:"_id,omitempty"`
	Deleted         bool      `bson:"deleted"`
	SagaTimestamp   int64     `bson:"saga_timestamp"`
	UserId          string    `bson:"userId,omitempty"`
	HostId          string    `bson:"hostId,omitempty"`
	AccommodationId string    `bson:"accommodationId,omitempty"`
	Rating          float64   `bson:"rating,omitempty"`
	TimeIssued      time.Time `bson:"timestamp,omitempty"`
}
