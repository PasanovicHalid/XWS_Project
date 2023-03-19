package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	Id       primitive.ObjectID
	Role     string
	Username string
	Password string
}
