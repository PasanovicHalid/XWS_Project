package persistance

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE                = "AccomodationDB"
	ACCOMODATION_COLLECTION = "Accomodations"
)

func NewMongoClient(host string, port string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	options := options.Client().ApplyURI(uri)
	return mongo.Connect(context.TODO(), options)
}
