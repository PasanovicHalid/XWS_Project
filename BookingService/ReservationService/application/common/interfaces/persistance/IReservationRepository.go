package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/domain"
)

type IReservationRepository interface {
	FindReservationById(ctx *context.Context, id string) (*domain.Reservation, error)
	CreateReservation(ctx *context.Context, reservation *domain.Reservation) error
	UpdateReservation(ctx *context.Context, reservation *domain.Reservation) error
	DeleteReservation(ctx *context.Context, id string) error
	GetAllReservations(ctx *context.Context) ([]*domain.Reservation, error)
	GetReservationsByAccommodationOfferID(ctx *context.Context, id string) ([]*domain.Reservation, error)
}
