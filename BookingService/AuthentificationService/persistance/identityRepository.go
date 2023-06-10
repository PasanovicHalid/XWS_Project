package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "AuthentificationDB"
	COLLECTION = "Identity"
)

type IdentityRepository struct {
	identities *mongo.Collection
	persistance.IIdentityRepository
}

func NewIdentityRepository(client *mongo.Client) *IdentityRepository {
	return &IdentityRepository{
		identities: client.Database(DATABASE).Collection(COLLECTION),
	}
}

func (repository *IdentityRepository) FindIdentityByUsername(ctx *context.Context, username string) (*domain.Identity, error) {
	filter := bson.D{{"username", username}, {"deleted", false}}
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
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objectId}, {"deleted", false}}
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
	result, err := repository.identities.InsertOne(*ctx, identity)
	if err != nil {
		return err
	}
	identity.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (repository *IdentityRepository) UpdateIdentity(ctx *context.Context, identity *domain.Identity) error {
	_, err := repository.identities.ReplaceOne(*ctx, bson.D{{"_id", identity.Id}}, identity)

	if err != nil {
		return err
	}

	return nil
}

func (repository *IdentityRepository) DeleteIdentity(ctx *context.Context, id string, sagaTimestamp int64) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	identity, err := repository.FindIdentityById(ctx, id)

	if err != nil {
		return "", err
	}

	result, err := repository.identities.UpdateOne(*ctx, bson.D{{"_id", objectId}, {"deleted", false}}, bson.D{{"$set", bson.D{{"deleted", true}, {"saga_timestamp", sagaTimestamp}}}})

	if err != nil {
		return "", err
	}

	if result.ModifiedCount == 0 {
		return "", persistance.ErrorIdentityNotFound
	}

	return identity.Role, nil
}

func (repository *IdentityRepository) RollbackDeleteIdentity(ctx *context.Context, id string, sagaTimestamp int64) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	result, err := repository.identities.UpdateOne(*ctx, bson.D{{"_id", objectId}, {"deleted", true}, {"saga_timestamp", sagaTimestamp}}, bson.D{{"$set", bson.D{{"deleted", false}, {"saga_timestamp", 0}}}})

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return persistance.ErrorIdentityNotFound
	}

	return nil
}

func (repository *IdentityRepository) CheckIfUsernameExists(ctx *context.Context, username string) (bool, error) {
	filter := bson.D{{"username", username}, {"deleted", false}}
	_, err := repository.filterOne(ctx, filter)

	if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return true, err
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
