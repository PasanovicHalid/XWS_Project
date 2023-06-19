package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/domain"
)

type IEmailRepository interface {
	UpdateWantedNotifications(ctx *context.Context, notifications *domain.WantedNotification) error
	SetWantedNotifications(ctx *context.Context, notifications *domain.WantedNotification) error
	// FindReservationById(ctx *context.Context, id string) (*domain.Reservation, error)
	// CreateReservation(ctx *context.Context, reservation *domain.Reservation) error
	// UpdateReservation(ctx *context.Context, reservation *domain.Reservation) error
	// DeleteReservation(ctx *context.Context, id string) error
	// DeleteReservationOfGuest(ctx *context.Context, id string, sagaTimestamp int64) error
	// ReverseDeleteReservationOfGuest(ctx *context.Context, id string, sagaTimestamp int64) error
	// DeleteReservationOfHost(ctx *context.Context, id string, sagaTimestamp int64) error
	// ReverseDeleteReservationOfHost(ctx *context.Context, id string, sagaTimestamp int64) error
	// GetAllReservations(ctx *context.Context) ([]*domain.Reservation, error)
	// GetReservationsByAccommodationOfferID(ctx *context.Context, id string) ([]*domain.Reservation, error)
}
