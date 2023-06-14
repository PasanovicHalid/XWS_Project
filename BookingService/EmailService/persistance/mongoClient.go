package persistance

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE         = "EmailDB"
	EMAIL_COLLECTION = "Emails"
)

func NewMongoClient(host string, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}
