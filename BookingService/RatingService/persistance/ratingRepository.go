package persistance

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingRepository struct {
	ratings *mongo.Collection
	persistance.IRatingRepository
}

func NewRatingRepository(client *mongo.Client) *RatingRepository {
	return &RatingRepository{
		ratings: client.Database(DATABASE).Collection(RATING_COLLECTION),
	}
}

func (repository *RatingRepository) GetAllRatingsMadeByCustomer(ctx *context.Context, id string) ([]*domain.Rating, error) {
	filter := bson.M{"userId": id, "deleted": false}
	return repository.filter(ctx, filter)
}

func (repository *RatingRepository) GetAllRatingsForHost(ctx *context.Context, id string) ([]*domain.Rating, error) {
	filter := bson.M{"hostId": id, "deleted": false}
	return repository.filter(ctx, filter)
}

func (repository *RatingRepository) CreateRating(ctx *context.Context, rating *domain.Rating) error {
	rating.TimeIssued = time.Now()

	_, err := repository.ratings.InsertOne(*ctx, rating)

	if err != nil {
		return err
	}

	return nil
}

func (repository *RatingRepository) UpdateRating(ctx *context.Context, id string, rating float64) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := repository.ratings.UpdateOne(*ctx, bson.M{"_id": objId, "deleted": false}, bson.M{"$set": bson.M{"rating": rating}})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (repository *RatingRepository) DeleteRating(ctx *context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := repository.ratings.UpdateOne(*ctx, bson.M{"_id": objId, "deleted": false}, bson.M{"$set": bson.M{"deleted": true}})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (repository *RatingRepository) DeleteAllRatingsMadeByCustomer(ctx *context.Context, id string) error {
	filter := bson.M{"userId": id, "deleted": false}
	update := bson.M{"$set": bson.M{"deleted": true}}

	_, err := repository.ratings.UpdateMany(*ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repository *RatingRepository) GetAllRatingsForAccommodation(ctx *context.Context, id string) ([]*domain.Rating, error) {
	filter := bson.M{"accommodationId": id, "deleted": false}
	return repository.filter(ctx, filter)
}

func (repository *RatingRepository) filter(ctx *context.Context, filter interface{}) ([]*domain.Rating, error) {
	cursor, err := repository.ratings.Find(*ctx, filter)
	defer cursor.Close(*ctx)

	if err != nil {
		return nil, err
	}

	return decode(ctx, cursor)
}

func (repository *RatingRepository) filterOne(ctx *context.Context, filter interface{}) (rating *domain.Rating, err error) {
	result := repository.ratings.FindOne(*ctx, filter)
	err = result.Decode(&rating)
	return
}

func decode(ctx *context.Context, cursor *mongo.Cursor) (ratings []*domain.Rating, err error) {
	for cursor.Next(*ctx) {
		var rating domain.Rating
		err := cursor.Decode(&rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, &rating)
	}
	err = cursor.Err()
	return
}
