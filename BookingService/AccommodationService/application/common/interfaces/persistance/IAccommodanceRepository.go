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
	UpdateAccommodationOffer(ctx *context.Context, reservation *domain.AccommodationOffer) error
	DeleteAccommodation(ctx *context.Context, id string) error
}
