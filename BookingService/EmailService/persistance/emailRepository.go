package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmailRepository struct {
	emails *mongo.Collection
}

func NewEmailRepository(client *mongo.Client) *EmailRepository {
	return &EmailRepository{
		emails: client.Database(DATABASE).Collection(EMAIL_COLLECTION),
	}
}

func (repository *EmailRepository) UpdateWantedNotifications(ctx *context.Context, notifications *domain.WantedNotification) error {
	filter := bson.M{"_id": notifications.UserId, "deleted": false}
	update := bson.M{"$set": notifications, "$setOnInsert": bson.M{"_id": notifications.UserId, "deleted": false}}

	opts := options.Update().SetUpsert(true)
	_, err := repository.emails.UpdateOne(*ctx, filter, update, opts)

	if err != nil {
		return err
	}

	return nil
}
