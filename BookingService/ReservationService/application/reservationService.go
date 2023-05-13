package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/domain"
)

type ReservationService struct {
	reservationRepository persistance.IReservationRepository
}

func NewReservationService(reservationRepository persistance.IReservationRepository) *ReservationService {
	return &ReservationService{
		reservationRepository: reservationRepository,
	}
}

func (service *ReservationService) GetReservationById(id string) (*domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.FindReservationById(&ctx, id)
}

func (service *ReservationService) CreateReservation(reservation *domain.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	return service.reservationRepository.CreateReservation(&ctx, reservation)
}

func (service *ReservationService) UpdateReservation(reservation *domain.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.UpdateReservation(&ctx, reservation)
}

func (service *ReservationService) DeleteReservation(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.DeleteReservation(&ctx, id)
}

func (service *ReservationService) GetAllReservations() ([]*domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.GetAllReservations(&ctx)
}

func (service *ReservationService) GetReservationByAccomodationId(id string) ([]*domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.GetReservationsByAccommodationOfferID(&ctx, id)
}
