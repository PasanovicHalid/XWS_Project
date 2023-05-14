package persistance

import (
	"context"
	"errors"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReservationRepository struct {
	reservations *mongo.Collection
}

func NewReservationRepository(client *mongo.Client) *ReservationRepository {
	return &ReservationRepository{
		reservations: client.Database(DATABASE).Collection(RESERVATION_COLLECTION),
	}
}

func (repository *ReservationRepository) FindReservationById(ctx *context.Context, id string) (*domain.Reservation, error) {
	filter := bson.D{{"_id", id}}
	result, err := repository.filterOne(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Reservation not found")

		}
		return nil, err
	}

	return result, nil
}

func (repository *ReservationRepository) CreateReservation(ctx *context.Context, reservation *domain.Reservation) error {
	_, err := repository.reservations.InsertOne(*ctx, reservation)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ReservationRepository) UpdateReservation(ctx *context.Context, reservation *domain.Reservation) error {
	_, err := repository.reservations.ReplaceOne(*ctx, bson.M{"_id": reservation.Id}, reservation)

	if err != nil {
		return err
	}

	return nil
}

func (repository *ReservationRepository) DeleteReservation(ctx *context.Context, id string) error {
	_, err := repository.reservations.DeleteOne(*ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}

	return nil
}

func (repository *ReservationRepository) filterOne(ctx *context.Context, filter interface{}) (reservation *domain.Reservation, err error) {
	result := repository.reservations.FindOne(*ctx, filter)
	err = result.Decode(&reservation)
	return
}

func (repository *ReservationRepository) GetAllReservations(ctx *context.Context) ([]*domain.Reservation, error) {
	filter := bson.D{} // Empty filter to retrieve all documents
	options := options.Find()

	cur, err := repository.reservations.Find(*ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)

	reservations := []*domain.Reservation{}
	for cur.Next(*ctx) {
		reservation := &domain.Reservation{}
		err := cur.Decode(&reservation)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}
func (repository *ReservationRepository) GetReservationsByAccommodationOfferID(ctx *context.Context, accommodationOfferID string) ([]*domain.Reservation, error) {
	filter := bson.M{"offerId": accommodationOfferID}
	options := options.Find()
	cur, err := repository.reservations.Find(*ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)

	reservations := []*domain.Reservation{}
	for cur.Next(*ctx) {
		reservation := &domain.Reservation{}
		err := cur.Decode(&reservation)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}
