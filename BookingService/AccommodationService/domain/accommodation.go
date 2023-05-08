package domain

type Accommodation struct {
	IdentityId       string `bson:"_id,omitempty"`
	Name             string `bson:"name,omitempty"`
	Location         string `bson:"location,omitempty"`
	Wifi             bool   `bson:"wifi,omitempty"`
	Kitchen          bool   `bson:"kitchen,omitempty"`
	AirConditioner   bool   `bson:"airConditioner,omitempty"`
	Parking          bool   `bson:"parking,omitempty"`
	MinNumberOfGuest int    `bson:"minGuest,omitempty"`
	MaxNumberOfGuest int    `bson:"maxGuest,omitempty"`
	Image            string `bson:"image,omitempty"`
}
