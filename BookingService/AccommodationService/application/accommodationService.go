package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/domain"
	accomodancePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/accommodation_service"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	"github.com/golang/protobuf/ptypes"
)

type AccommodationService struct {
	accomodationRepository persistance.IAccommodationRepository
}

func NewAccomodationService(accomodationRepository persistance.IAccommodationRepository) *AccommodationService {
	return &AccommodationService{
		accomodationRepository: accomodationRepository,
	}
}

func (service *AccommodationService) CreateAccomodation(newAccomodation *accomodancePB.NewAccomodation) (*common_pb.RequestResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	newAccomodationClass := mapNewAccomodationToAccommodation(newAccomodation)
	err := service.accomodationRepository.CreateAccomodation(&ctx, newAccomodationClass)
	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEUSPESNO",
		}, nil
	}
	return &common_pb.RequestResult{
		Code:    200,
		Message: "USPESNO",
	}, nil
}

func mapNewAccomodationToAccommodation(newAccomodation *accomodancePB.NewAccomodation) *domain.Accommodation {
	return &domain.Accommodation{
		Id:               newAccomodation.GetId(),
		Name:             newAccomodation.GetName(),
		OwnerId:          newAccomodation.GetOwnerId(),
		Location:         newAccomodation.GetLocation(),
		Wifi:             newAccomodation.GetWifi(),
		Kitchen:          newAccomodation.GetKitchen(),
		AirConditioner:   newAccomodation.GetAirConditioner(),
		Parking:          newAccomodation.GetParking(),
		MinNumberOfGuest: int(newAccomodation.GetMinNumberOfGuests()),
		MaxNumberOfGuest: int(newAccomodation.GetMaxNumberOfGuests()),
		Images:           newAccomodation.GetImages(),
	}
}

func (service *AccommodationService) CreateAccomodationOffer(newAccomodationOffer *accomodancePB.CreateOfferRequest) (*common_pb.RequestResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	accomodationOffer, err := CreateOfferRequestToAccommodationOffer(newAccomodationOffer)
	if service.DateOverlapCheck(accomodationOffer) {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEISPRAVNO UNETI PODACI",
		}, nil
	}
	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEUSPESNO UNETI PODACI",
		}, nil
	} else {
		err2 := service.accomodationRepository.CreateAccomodationOffer(&ctx, accomodationOffer)
		if err2 != nil {
			return &common_pb.RequestResult{
				Code:    500,
				Message: "NEUSPESNO CUVANJE",
			}, nil
		}

		return &common_pb.RequestResult{
			Code:    200,
			Message: "USPESNO",
		}, nil
	}
}

func (service *AccommodationService) UpdateAccommodationOffer(newAccomodationOffer *accomodancePB.AccommodationOffer) (*common_pb.RequestResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	accomodationOffer, err := CreateOfferRequestToAccommodationOfferWithId(newAccomodationOffer)
	if service.DateOverlapCheckWithId(accomodationOffer) {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEISPRAVNO UNETI PODACI",
		}, nil
	}
	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: "NEUSPESNO UNETI PODACI",
		}, nil
	} else {
		err2 := service.accomodationRepository.UpdateAccommodationOffer(&ctx, accomodationOffer)
		if err2 != nil {
			return &common_pb.RequestResult{
				Code:    500,
				Message: err2.Error(),
			}, nil
		}

		return &common_pb.RequestResult{
			Code:    200,
			Message: "USPESNO",
		}, nil
	}
}

func CreateOfferRequestToAccommodationOffer(req *accomodancePB.CreateOfferRequest) (*domain.AccommodationOffer, error) {
	startDateTimeUTC, err := ptypes.Timestamp(req.StartDateTimeUtc)
	if err != nil {
		return nil, err
	}

	endDateTimeUTC, err := ptypes.Timestamp(req.EndDateTimeUtc)
	if err != nil {
		return nil, err
	}

	return &domain.AccommodationOffer{
		AccommodationId:           req.AccommodationId,
		AvailableStartDateTimeUTC: startDateTimeUTC,
		AvailableEndDateTimeUTC:   endDateTimeUTC,
		Price:                     int(req.GetPrice()),
		PerGuest:                  req.PerGuest,
	}, nil
}

func CreateOfferRequestToAccommodationOfferWithId(req *accomodancePB.AccommodationOffer) (*domain.AccommodationOffer, error) {
	startDateTimeUTC, err := ptypes.Timestamp(req.StartDateTimeUtc)
	if err != nil {
		return nil, err
	}

	endDateTimeUTC, err := ptypes.Timestamp(req.EndDateTimeUtc)
	if err != nil {
		return nil, err
	}

	return &domain.AccommodationOffer{
		Id:                        req.Id,
		AccommodationId:           req.AccommodationId,
		AvailableStartDateTimeUTC: startDateTimeUTC,
		AvailableEndDateTimeUTC:   endDateTimeUTC,
		Price:                     int(req.GetPrice()),
		PerGuest:                  req.PerGuest,
	}, nil
}

func (service *AccommodationService) DateOverlapCheck(offer *domain.AccommodationOffer) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	offers, _ := service.accomodationRepository.GetAllAccommodationOffers(&ctx)
	for _, o := range offers {
		if o.AccommodationId == offer.AccommodationId {
			// Case 1: The new offer's start date is between the start and end date of an existing offer
			if offer.AvailableStartDateTimeUTC.After(o.AvailableStartDateTimeUTC) && offer.AvailableStartDateTimeUTC.Before(o.AvailableEndDateTimeUTC) {
				return true
			}
			// Case 2: The new offer's end date is between the start and end date of an existing offer
			if offer.AvailableEndDateTimeUTC.After(o.AvailableStartDateTimeUTC) && offer.AvailableEndDateTimeUTC.Before(o.AvailableEndDateTimeUTC) {
				return true
			}
			// Case 3: The new offer completely covers an existing offer
			if offer.AvailableStartDateTimeUTC.Before(o.AvailableStartDateTimeUTC) && offer.AvailableEndDateTimeUTC.After(o.AvailableEndDateTimeUTC) {
				return true
			}
			// Case 4: An existing offer completely covers the new offer
			if offer.AvailableStartDateTimeUTC.After(o.AvailableStartDateTimeUTC) && offer.AvailableEndDateTimeUTC.Before(o.AvailableEndDateTimeUTC) {
				return true
			}
		}
	}
	return false
}

func (service *AccommodationService) DateOverlapCheckWithId(offer *domain.AccommodationOffer) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	offers, _ := service.accomodationRepository.GetAllAccommodationOffers(&ctx)
	for _, o := range offers {
		if o.Id != offer.Id && o.AccommodationId == offer.AccommodationId {
			if offer.AvailableStartDateTimeUTC.After(o.AvailableStartDateTimeUTC) && offer.AvailableStartDateTimeUTC.Before(o.AvailableEndDateTimeUTC) {
				return true
			}
			// Case 2: The new offer's end date is between the start and end date of an existing offer
			if offer.AvailableEndDateTimeUTC.After(o.AvailableStartDateTimeUTC) && offer.AvailableEndDateTimeUTC.Before(o.AvailableEndDateTimeUTC) {
				return true
			}
			// Case 3: The new offer completely covers an existing offer
			if offer.AvailableStartDateTimeUTC.Before(o.AvailableStartDateTimeUTC) && offer.AvailableEndDateTimeUTC.After(o.AvailableEndDateTimeUTC) {
				return true
			}
			// Case 4: An existing offer completely covers the new offer
			if offer.AvailableStartDateTimeUTC.After(o.AvailableStartDateTimeUTC) && offer.AvailableEndDateTimeUTC.Before(o.AvailableEndDateTimeUTC) {
				return true
			}
		}
		// Case 1: The new offer's start date is between the start and end date of an existing offer
	}
	return false
}

func (service *AccommodationService) GetAllAccommodationsByOwner(identityId string) []domain.Accommodation {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	filteredAccommodations := []domain.Accommodation{}
	accomodations, _ := service.accomodationRepository.GetAllAccommodations(&ctx)
	for _, accomodation := range accomodations {
		if accomodation.OwnerId == identityId {
			filteredAccommodations = append(filteredAccommodations, *accomodation)
		}
	}
	return filteredAccommodations
}

func (service *AccommodationService) DeleteAllAccommodationsByOwner(identityId string) (*common_pb.RequestResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	for _, accommodation := range service.GetAllAccommodationsByOwner(identityId) {
		err := service.accomodationRepository.DeleteAccommodation(&ctx, accommodation.Id)
		if err != nil {
			return &common_pb.RequestResult{
				Code:    500,
				Message: "BRISANJE NEUSPESNO",
			}, err
		}
	}
	return &common_pb.RequestResult{
		Code:    200,
		Message: "BRISANJE NEUSPESNO",
	}, nil
}

func (service *AccommodationService) FilterAccommodations(message *accomodancePB.AccommodationSearch) (*accomodancePB.GetFilteredAccommodationsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	var filteredAccommodations []*accomodancePB.NewAccomodation
	accommodations, _ := service.accomodationRepository.GetAllAccommodations(&ctx)
	accommodationOffers, _ := service.accomodationRepository.GetAllAccommodationOffers(&ctx)
	for _, accommodation := range accommodations {
		for _, offer := range accommodationOffers {
			if offer.AccommodationId == accommodation.Id {
				if offer.AvailableStartDateTimeUTC.Before(message.StartDateTimeUtc.AsTime()) &&
					offer.AvailableEndDateTimeUTC.After(message.EndDateTimeUtc.AsTime()) &&
					accommodation.Location == message.Location &&
					accommodation.MinNumberOfGuest <= int(message.GuestNumber) &&
					accommodation.MaxNumberOfGuest >= int(message.GuestNumber) {
					filteredAccommodations = append(filteredAccommodations, convertToNewAccomodation(*accommodation))
					break
				}
			}
		}
	}
	response := &accomodancePB.GetFilteredAccommodationsResponse{
		FilteredAccommodations: filteredAccommodations,
	}

	return response, nil
}

func convertToNewAccomodation(accommodation domain.Accommodation) *accomodancePB.NewAccomodation {
	return &accomodancePB.NewAccomodation{
		Id:                accommodation.Id,
		Name:              accommodation.Name,
		Location:          accommodation.Location,
		Wifi:              accommodation.Wifi,
		Kitchen:           accommodation.Kitchen,
		AirConditioner:    accommodation.AirConditioner,
		Parking:           accommodation.Parking,
		MinNumberOfGuests: int32(accommodation.MinNumberOfGuest),
		MaxNumberOfGuests: int32(accommodation.MaxNumberOfGuest),
		Images:            accommodation.Images,
		OwnerId:           accommodation.OwnerId,
	}
}
