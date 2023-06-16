package persistance

import (
	"context"
	"fmt"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	_, err := repository.emails.ReplaceOne(*ctx, bson.M{"_id": notifications.UserId, "deleted": false}, notifications)

	if err != nil {
		return err
	}

	return nil
}

func (repository *EmailRepository) SetWantedNotifications(ctx *context.Context, notifications *domain.WantedNotification) error {
	fmt.Print("OOOOO\n")
	_, err := repository.emails.InsertOne(*ctx, notifications)
	fmt.Print(err)
	fmt.Print("OLOLOL\n\n")
	if err != nil {
		return err
	}

	return nil
}
