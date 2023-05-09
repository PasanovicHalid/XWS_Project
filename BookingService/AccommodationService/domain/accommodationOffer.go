package domain

type AccommodationOffer struct {
	Id              string    `bson:"_id,omitempty"`
	AccommodationId string    `bson:"accommodationId,omitempty"`
	DateRange       DateRange `bson:"dateRange,omitempty"`
	Price           int       `bson:"price,omitempty"`
}
