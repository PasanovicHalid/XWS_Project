package application

import (
	"fmt"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/persistance"
)

type AccommodationService struct {
	accomodationRepository persistance.AccommodationRepository
}

func NewAccomodationService(accomodationRepository persistance.AccommodationRepository) *AccommodationService {
	return &AccommodationService{
		accomodationRepository: accomodationRepository,
	}
}

func (s *AccommodationService) PrintSuccess(message string) {
	fmt.Printf("Success: %s\n", message)
}
