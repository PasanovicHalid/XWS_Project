package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Identity struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role           string             `bson:"role" json:"role"`
	Username       string             `bson:"username" json:"username"`
	Password       string             `bson:"password" json:"password"`
	ApiKey         string             `bson:"apiKey" json:"apiKey"`
	ApiKeyDuration time.Time          `bson:"apiKeyDuration,omitempty" json:"apiKeyDuration,omitempty"`
}
