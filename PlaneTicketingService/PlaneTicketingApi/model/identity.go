package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Role     string             `bson:"role" json:"role"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}
