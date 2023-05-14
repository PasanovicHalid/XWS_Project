package domain

type Accommodation struct {
	Id               string   `bson:"_id,omitempty"`
	Name             string   `bson:"name,omitempty"`
	OwnerId          string   `bson:"ownerId,omitempty"`
	Location         string   `bson:"location,omitempty"`
	Wifi             bool     `bson:"wifi,omitempty"`
	Kitchen          bool     `bson:"kitchen,omitempty"`
	AirConditioner   bool     `bson:"airConditioner,omitempty"`
	Parking          bool     `bson:"parking,omitempty"`
	MinNumberOfGuest int      `bson:"minGuest,omitempty"`
	MaxNumberOfGuest int      `bson:"maxGuest,omitempty"`
	Images           []string `bson:"images,omitempty"`
}
