package model

type User struct {
	Identity  Identity `bson:"identity,inline" json:"identity"`
	Firstname string   `bson:"firstname" json:"firstname"`
	Lastname  string   `bson:"lastname" json:"lastname"`
}
