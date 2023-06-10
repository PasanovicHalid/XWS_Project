package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Identity struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Deleted       bool               `bson:"deleted"`
	SagaTimestamp int64              `bson:"saga_timestamp"`
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
	Role          string             `bson:"role"`
}
