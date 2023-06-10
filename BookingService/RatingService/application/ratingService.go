package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/domain"
)

type RatingService struct {
	ratingRepository persistance.IRatingRepository
}

func NewRatingService(ratingRepository persistance.IRatingRepository) *RatingService {
	return &RatingService{ratingRepository: ratingRepository}
}

func (service *RatingService) GetAllRatingsMadeByCustomer(id string) ([]*domain.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.GetAllRatingsMadeByCustomer(&ctx, id)
}

func (service *RatingService) GetAllRatingsForHost(id string) ([]*domain.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.GetAllRatingsForHost(&ctx, id)
}

func (service *RatingService) CreateRating(rating *domain.Rating) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.CreateRating(&ctx, rating)
}

func (service *RatingService) UpdateRating(id string, rating float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.UpdateRating(&ctx, id, rating)
}

func (service *RatingService) DeleteRating(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.DeleteRating(&ctx, id)
}
