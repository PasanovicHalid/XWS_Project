package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccommodationRepository struct {
	accomodations        *mongo.Collection
	accommodationsOffers *mongo.Collection
}

const (
	DATABASE                        = "AccommodationDB"
	ACCOMMODATION_COLLECTION        = "Accommodations"
	ACCOMMODATION_OFFERS_COLLECTION = "AccommodationOffers"
)

func NewAccommodationRepository(client *mongo.Client) *AccommodationRepository {
	return &AccommodationRepository{
		accomodations:        client.Database(DATABASE).Collection(ACCOMMODATION_COLLECTION),
		accommodationsOffers: client.Database(DATABASE).Collection(ACCOMMODATION_OFFERS_COLLECTION),
	}
}

func (repository *AccommodationRepository) CreateAccomodation(ctx *context.Context, reservation *domain.Accommodation) error {
	_, err := repository.accomodations.InsertOne(*ctx, reservation)
	return err
}

func (repository *AccommodationRepository) CreateAccomodationOffer(ctx *context.Context, reservation *domain.AccommodationOffer) error {
	_, err := repository.accommodationsOffers.InsertOne(*ctx, reservation)
	return err
}

func (repository *AccommodationRepository) GetAllAccommodationOffers(ctx *context.Context) ([]*domain.AccommodationOffer, error) {
	filter := bson.D{}
	options := options.Find()

	cur, err := repository.accommodationsOffers.Find(*ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)

	offers := []*domain.AccommodationOffer{}
	for cur.Next(*ctx) {
		reservation := &domain.AccommodationOffer{}
		err := cur.Decode(&reservation)
		if err != nil {
			return nil, err
		}
		offers = append(offers, reservation)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return offers, nil
}

func (repository *AccommodationRepository) UpdateAccommodationOffer(ctx *context.Context, reservation *domain.AccommodationOffer) error {
	id, _ := primitive.ObjectIDFromHex(reservation.Id)
	temp := &domain.AccommodationOffer{
		AccommodationId:           reservation.AccommodationId,
		AvailableStartDateTimeUTC: reservation.AvailableEndDateTimeUTC,
		AvailableEndDateTimeUTC:   reservation.AvailableEndDateTimeUTC,
		Price:                     reservation.Price,
		PerGuest:                  reservation.PerGuest,
		AutomaticAcceptation:      reservation.AutomaticAcceptation,
	}
	_, err := repository.accommodationsOffers.ReplaceOne(*ctx, bson.M{"_id": id}, temp)
	if err != nil {
		return err
	}

	return nil
}

func (repository *AccommodationRepository) GetAllAccommodations(ctx *context.Context) ([]*domain.Accommodation, error) {
	filter := bson.D{}
	options := options.Find()

	cur, err := repository.accomodations.Find(*ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*ctx)

	offers := []*domain.Accommodation{}
	for cur.Next(*ctx) {
		reservation := &domain.Accommodation{}
		err := cur.Decode(&reservation)
		if err != nil {
			return nil, err
		}
		offers = append(offers, reservation)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return offers, nil
}

func (repository *AccommodationRepository) DeleteAccommodation(ctx *context.Context, id string) error {
	_, err := repository.accomodations.DeleteOne(*ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}

	return nil
}

func (repository *AccommodationRepository) GetAccommodationById(ctx *context.Context, id string) (*domain.Accommodation, error) {
	objID := id

	filter := bson.M{"_id": objID}
	result := repository.accomodations.FindOne(*ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	accommodation := &domain.Accommodation{}
	err := result.Decode(accommodation)
	if err != nil {
		return nil, err
	}

	return accommodation, nil
}

func (repository *AccommodationRepository) GetAccommodationOfferById(ctx *context.Context, id string) (*domain.AccommodationOffer, error) {
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID}
	result := repository.accommodationsOffers.FindOne(*ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	accommodationOffer := &domain.AccommodationOffer{}
	err := result.Decode(accommodationOffer)
	if err != nil {
		return nil, err
	}

	return accommodationOffer, nil
}
