package domain

type User struct {
	IdentityId  string `bson:"_id,omitempty"`
	FirstName   string `bson:"firstName,omitempty"`
	LastName    string `bson:"lastName,omitempty"`
	Email       string `bson:"email,omitempty"`
	PhoneNumber string `bson:"phoneNumber,omitempty"`
	Address     string `bson:"address,omitempty"`
}
