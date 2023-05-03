package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	KEY_DATABASE   = "AuthentificationDB"
	KEY_COLLECTION = "Key"
)

type KeyRepository struct {
	keys *mongo.Collection
}

func NewKeyRepository(client *mongo.Client) *KeyRepository {
	return &KeyRepository{
		keys: client.Database(KEY_DATABASE).Collection(KEY_COLLECTION),
	}
}

func (repository *KeyRepository) GetKey(ctx *context.Context) (*domain.Key, error) {
	filter := bson.D{{}}
	return repository.filterOne(ctx, filter)
}

func (repository *KeyRepository) SaveKey(ctx *context.Context, key *domain.Key) error {
	_, err := repository.keys.DeleteMany(*ctx, bson.D{{}})
	if err != nil {
		return err
	}

	_, err = repository.keys.InsertOne(*ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (repository *KeyRepository) filterOne(ctx *context.Context, filter interface{}) (Key *domain.Key, err error) {
	result := repository.keys.FindOne(*ctx, filter)
	err = result.Decode(&Key)
	return
}
