package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseMongoDb struct {
	client *mongo.Client
	logger *log.Logger
}

type DatabaseCollection struct {
	Collection *mongo.Collection
	Logger     *log.Logger
}

var MongoInstance *DatabaseMongoDb

func NewDb(ctx context.Context, logger *log.Logger) (*DatabaseMongoDb, error) {
	dburi := os.Getenv("MONGO_DB_URI")
	if len(dburi) == 0 {
		dburi = "mongodb://localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &DatabaseMongoDb{
		client: client,
		logger: logger,
	}, nil
}

func (db *DatabaseMongoDb) Disconnect(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DatabaseMongoDb) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := db.client.Ping(ctx, readpref.Primary())
	if err != nil {
		db.logger.Println(err)
	}

	// Print available databases
	databases, err := db.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		db.logger.Println(err)
	}
	fmt.Println(databases)
}

func OpenCollection(db *DatabaseMongoDb, collectionName string) *DatabaseCollection {
	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		dbName = "TicketingDB"
	}

	collection := db.client.Database(dbName).Collection(collectionName)

	return &DatabaseCollection{
		Collection: collection,
		Logger:     db.logger,
	}
}

func SetupDb(timeoutContext context.Context, storeLogger *log.Logger, logger *log.Logger) *DatabaseMongoDb {
	db, err := NewDb(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}

	db.Ping()

	return db
}
