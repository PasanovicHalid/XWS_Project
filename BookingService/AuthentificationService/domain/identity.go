package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}
