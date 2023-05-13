package application

import (
	"fmt"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application/common/interfaces/persistance"
)

type AccommodationService struct {
	accomodationRepository persistance.IAccommodationRepository
}

func NewAccomodationService(accomodationRepository persistance.IAccommodationRepository) *AccommodationService {
	return &AccommodationService{
		accomodationRepository: accomodationRepository,
	}
}

func (s *AccommodationService) PrintSuccess(message string) {
	fmt.Printf("Success: %s\n", message)
}

// func (s *AccommodationService) TempServiceMethod(message *accomodancePB.TempMessage) (*common_pb.RequestResult, error) {
// 	fmt.Println("USPEH")
// 	return &common_pb.RequestResult{
// 		Code:    200,
// 		Message: "USPESNO",
// 	}, nil
// }
