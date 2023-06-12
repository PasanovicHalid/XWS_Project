package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/domain"
)

type IAccommodationRepository interface {
	CreateAccomodation(ctx *context.Context, reservation *domain.Accommodation) error
	CreateAccomodationOffer(ctx *context.Context, reservation *domain.AccommodationOffer) error
	GetAllAccommodationOffers(ctx *context.Context) ([]*domain.AccommodationOffer, error)
	GetAllAccommodations(ctx *context.Context) ([]*domain.Accommodation, error)
	GetAllAccommodationsByIdList(ctx *context.Context, idList []string) ([]*domain.Accommodation, error)
	GetAllAccommodationsSaga(ctx *context.Context, sagaTimestamp int64) ([]*domain.Accommodation, error)
	UpdateAccommodationOffer(ctx *context.Context, reservation *domain.AccommodationOffer) error
	DeleteAccommodation(ctx *context.Context, id string, sagaTimestamp int64) error
	ReverseDeleteAccommodation(ctx *context.Context, id string, sagaTimestamp int64) error
	DeleteAccommodationOffers(ctx *context.Context, id string, sagaTimestamp int64) error
	ReverseDeleteAccommodationOffers(ctx *context.Context, id string, sagaTimestamp int64) error
	GetAccommodationById(ctx *context.Context, id string) (*domain.Accommodation, error)
	GetAccommodationOfferById(ctx *context.Context, id string) (*domain.AccommodationOffer, error)
}
