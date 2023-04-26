package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "AuthentificationDB"
	COLLECTION = "Identity"
)

type IdentityRepository struct {
	identities *mongo.Collection
}

func NewIdentityRepository(client *mongo.Client) *IdentityRepository {
	return &IdentityRepository{
		identities: client.Database(DATABASE).Collection(COLLECTION),
	}
}

func (repository *IdentityRepository) FindIdentityByEmail(ctx *context.Context, email string) (*domain.Identity, error) {
	filter := bson.D{{"email", email}}
	result, err := repository.filterOne(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, persistance.ErrorIdentityNotFound
		}
		return nil, err
	}

	return result, nil
}

func (repository *IdentityRepository) FindIdentityById(ctx *context.Context, id string) (*domain.Identity, error) {
	filter := bson.D{{"_id", id}}
	result, err := repository.filterOne(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, persistance.ErrorIdentityNotFound
		}
		return nil, err
	}

	return result, nil
}

func (repository *IdentityRepository) InsertIdentity(ctx *context.Context, identity *domain.Identity) error {
	_, err := repository.identities.InsertOne(*ctx, identity)
	if err != nil {
		return err
	}
	return nil
}

func (repository *IdentityRepository) UpdateIdentity(ctx *context.Context, identity *domain.Identity) error {
	_, err := repository.identities.ReplaceOne(*ctx, bson.D{{"_id", identity.Id}}, identity)

	if err != nil {
		return err
	}

	return nil
}

func (repository *IdentityRepository) DeleteIdentity(ctx *context.Context, id string) error {
	result, err := repository.identities.DeleteOne(*ctx, bson.D{{"_id", id}})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return persistance.ErrorIdentityNotFound
	}

	return nil
}

func (repository *IdentityRepository) CheckIfEmailExists(ctx *context.Context, email string) (bool, error) {
	filter := bson.D{{"email", email}}
	_, err := repository.filterOne(ctx, filter)

	if err == mongo.ErrNoDocuments {
		return true, nil
	}

	return false, err
}

func (repository *IdentityRepository) filter(ctx *context.Context, filter interface{}) ([]*domain.Identity, error) {
	cursor, err := repository.identities.Find(*ctx, filter)
	defer cursor.Close(*ctx)

	if err != nil {
		return nil, err
	}

	return decode(ctx, cursor)
}

func (repository *IdentityRepository) filterOne(ctx *context.Context, filter interface{}) (identity *domain.Identity, err error) {
	result := repository.identities.FindOne(*ctx, filter)
	err = result.Decode(&identity)
	return
}

func decode(ctx *context.Context, cursor *mongo.Cursor) (identities []*domain.Identity, err error) {
	for cursor.Next(*ctx) {
		var identity domain.Identity
		err := cursor.Decode(&identity)
		if err != nil {
			return nil, err
		}
		identities = append(identities, &identity)
	}
	err = cursor.Err()
	return
}
