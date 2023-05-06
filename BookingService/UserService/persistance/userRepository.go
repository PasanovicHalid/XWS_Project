package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	users *mongo.Collection
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		users: client.Database(DATABASE).Collection(USER_COLLECTION),
	}
}

func (repository *UserRepository) FindUserById(ctx *context.Context, id string) (*domain.User, error) {
	filter := bson.D{{"_id", id}}
	result, err := repository.filterOne(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, persistance.ErrorUserNotFound
		}
		return nil, err
	}

	return result, nil
}

func (repository *UserRepository) CreateUser(ctx *context.Context, user *domain.User) error {
	_, err := repository.users.InsertOne(*ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) UpdateUser(ctx *context.Context, user *domain.User) error {
	_, err := repository.users.ReplaceOne(*ctx, bson.M{"_id": user.IdentityId}, user)

	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) filterOne(ctx *context.Context, filter interface{}) (user *domain.User, err error) {
	result := repository.users.FindOne(*ctx, filter)
	err = result.Decode(&user)
	return
}