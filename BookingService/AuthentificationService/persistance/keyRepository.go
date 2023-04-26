package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
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

func (repository *KeyRepository) GetKeyPair(ctx *context.Context) (*domain.KeyPair, error) {
	filter := bson.D{{}}
	result, err := repository.filterOne(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, persistance.ErrorKeyNotFound
		}
		return nil, err
	}

	return result, nil
}

func (repository *KeyRepository) InsertNewOrReplaceOldKeyPair(ctx *context.Context, keyPair *domain.KeyPair) error {
	_, err := repository.keys.DeleteMany(*ctx, bson.D{{}})
	if err != nil {
		return err
	}

	_, err = repository.keys.InsertOne(*ctx, keyPair)
	if err != nil {
		return err
	}

	return nil
}

func (repository *KeyRepository) filterOne(ctx *context.Context, filter interface{}) (KeyPair *domain.KeyPair, err error) {
	result := repository.keys.FindOne(*ctx, filter)
	err = result.Decode(&KeyPair)
	return
}
