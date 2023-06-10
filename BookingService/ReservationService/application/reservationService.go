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

func (service *ReservationService) DeleteReservationOfGuest(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.DeleteReservationOfGuest(&ctx, id, sagaTimestamp)
}

func (service *ReservationService) ReverseDeleteReservationOfGuest(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.ReverseDeleteReservationOfGuest(&ctx, id, sagaTimestamp)
}

func (service *ReservationService) DeleteReservationOfHost(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.DeleteReservationOfHost(&ctx, id, sagaTimestamp)
}

func (service *ReservationService) ReverseDeleteReservationOfHost(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.ReverseDeleteReservationOfHost(&ctx, id, sagaTimestamp)
}

func (service *ReservationService) CheckHostActiveReservaton(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	reservations, err := service.reservationRepository.GetAllReservations(&ctx)
	if err != nil {
		return false, err
	}

	hasActiveReservations := false
	for _, reservation := range reservations {
		if reservation.HostId == id && reservation.ReservationStatus == domain.Accepted {
			hasActiveReservations = true
			break
		}
	}

	return hasActiveReservations, nil
}

func (service *ReservationService) CheckGuestActiveReservaton(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	reservations, err := service.reservationRepository.GetAllReservations(&ctx)
	if err != nil {
		return false, err
	}

	hasActiveReservations := false
	for _, reservation := range reservations {
		if reservation.CustomerId == id && reservation.ReservationStatus == domain.Accepted {
			hasActiveReservations = true
			break
		}
	}

	return hasActiveReservations, nil
}

func (service *ReservationService) GetAllReservations() ([]*domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.GetAllReservations(&ctx)
}

func (service *ReservationService) GetReservationsByAccommodationOfferID(id string) ([]*domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.reservationRepository.GetReservationsByAccommodationOfferID(&ctx, id)
}
