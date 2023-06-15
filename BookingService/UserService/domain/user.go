package domain

type User struct {
	IdentityId    string `bson:"_id,omitempty"`
	Deleted       bool   `bson:"deleted"`
	SagaTimestamp int64  `bson:"saga_timestamp"`
	FirstName     string `bson:"firstName,omitempty"`
	LastName      string `bson:"lastName,omitempty"`
	Email         string `bson:"email,omitempty"`
	PhoneNumber   string `bson:"phoneNumber,omitempty"`
	Address       string `bson:"address,omitempty"`
	Distinguished bool   `bson:"distinguished,omitempty"`
}
