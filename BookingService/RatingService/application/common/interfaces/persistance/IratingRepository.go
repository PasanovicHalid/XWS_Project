package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/domain"
)

type IRatingRepository interface {
	GetAllRatingsMadeByCustomer(ctx *context.Context, id string) ([]*domain.Rating, error)
	GetAllRatingsForHost(ctx *context.Context, id string) ([]*domain.Rating, error)
	GetAllRatingsForAccommodation(ctx *context.Context, id string) ([]*domain.Rating, error)
	CreateRating(ctx *context.Context, rating *domain.Rating) error
	UpdateRating(ctx *context.Context, id string, rating float64) error
	DeleteRating(ctx *context.Context, id string) error
}
