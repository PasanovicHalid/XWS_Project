package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application/common/interfaces/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/domain"
)

type RatingService struct {
	ratingRepository   persistance.IRatingRepository
	notificationSender message_queues.INotificationSender
}

func NewRatingService(ratingRepository persistance.IRatingRepository, notificationSender message_queues.INotificationSender) *RatingService {
	return &RatingService{
		ratingRepository:   ratingRepository,
		notificationSender: notificationSender,
	}
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

func (service *RatingService) DeleteAllRatingsMadeByCustomer(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.DeleteAllRatingsMadeByCustomer(&ctx, id)
}

func (service *RatingService) GetAllRatingsForAccommodation(id string) ([]*domain.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.ratingRepository.GetAllRatingsForAccommodation(&ctx, id)
}

func (service *RatingService) GetAverageRatingForHost(id string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	ratings, err := service.ratingRepository.GetAllRatingsForHost(&ctx, id)

	if err != nil {
		return 0, err
	}

	var sum float64 = 0
	for _, rating := range ratings {
		sum += rating.Rating
	}

	if len(ratings) == 0 {
		return 0, nil
	}

	return sum / float64(len(ratings)), nil
}

func (service *RatingService) GetRatingForAccommodation(id string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	ratings, err := service.ratingRepository.GetAllRatingsForAccommodation(&ctx, id)

	if err != nil {
		return 0, err
	}

	var sum float64 = 0
	for _, rating := range ratings {
		sum += rating.Rating
	}

	if len(ratings) == 0 {
		return 0, nil
	}

	return sum / float64(len(ratings)), nil
}
