package presentation

import (
	"context"
	"fmt"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/domain"
	accomodancePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/accommodation_service"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
)

type AccommodationHandler struct {
	accomodationService *application.AccommodationService
	accomodancePB.UnimplementedAccommodationServiceServer
}

func NewAccomodationHandler(accomodationService *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		accomodationService: accomodationService,
	}
}

func (handler *AccommodationHandler) CreateAccomodation(ctx context.Context, message *accomodancePB.NewAccomodation) (*common_pb.RequestResult, error) {
	return handler.accomodationService.CreateAccomodation(message)
}

func (handler *AccommodationHandler) CreateAccomodationOffer(ctx context.Context, message *accomodancePB.CreateOfferRequest) (*common_pb.RequestResult, error) {
	return handler.accomodationService.CreateAccomodationOffer(message)
}

func (handler *AccommodationHandler) UpdateAccomodationOffer(ctx context.Context, message *accomodancePB.AccommodationOffer) (*common_pb.RequestResult, error) {
	fmt.Println(message.GetStartDateTimeUtc().String())
	fmt.Println(message.GetEndDateTimeUtc().String())
	return handler.accomodationService.UpdateAccommodationOffer(message)
}

func (handler *AccommodationHandler) FilterAccommodations(ctx context.Context, message *accomodancePB.AccommodationSearch) (*accomodancePB.GetFilteredAccommodationsResponse, error) {
	return handler.accomodationService.FilterAccommodations(message)
}

func (handler *AccommodationHandler) GetOwnerIdByAccommodationId(ctx context.Context, message *accomodancePB.GetOwnerIdRequest) (*accomodancePB.GetOwnerIdResponse, error) {
	return handler.accomodationService.GetOwnerIdByAccommodationId(message)
}

func (handler *AccommodationHandler) SetAutomaticAcception(ctx context.Context, message *accomodancePB.SetAutomaticStatusRequest) (*accomodancePB.SetAutomaticStatusResponse, error) {
	return handler.accomodationService.SetAutomaticAcception(message)
}

func (handler *AccommodationHandler) GetAutomaticAcception(ctx context.Context, message *accomodancePB.GetAutomaticStatusRequest) (*accomodancePB.GetAutomaticStatusResponse, error) {
	return handler.accomodationService.GetAutomaticAcception(message)
}

func (handler *AccommodationHandler) GetAllAccommodationsByOwner(ctx context.Context, id *accomodancePB.IdentityIdRequest) (*accomodancePB.GetFilteredAccommodationsResponse, error) {
	accommodations := handler.accomodationService.GetAllAccommodationsByOwner(id.GetId())
	return ConvertToGetFilteredAccommodationsResponse(&accommodations)
}

func (handler *AccommodationHandler) GetAllAccommodationsByIdList(ctx context.Context, request *accomodancePB.IdListRequest) (*accomodancePB.GetFilteredAccommodationsResponse, error) {
	accommodations, err := handler.accomodationService.GetAllAccommodationsByIdList(request.Ids)

	if err != nil {
		return nil, err
	}

	response := &accomodancePB.GetFilteredAccommodationsResponse{}
	for _, accommodation := range accommodations {
		newAccommodation := convertToNewAccommodation(*accommodation)
		response.FilteredAccommodations = append(response.FilteredAccommodations, newAccommodation)
	}
	return response, nil
}

func convertToNewAccommodation(accommodation domain.Accommodation) *accomodancePB.NewAccomodation {
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

func ConvertToGetFilteredAccommodationsResponse(accommodations *[]domain.Accommodation) (*accomodancePB.GetFilteredAccommodationsResponse, error) {
	response := &accomodancePB.GetFilteredAccommodationsResponse{}
	for _, accommodation := range *accommodations {
		newAccommodation := convertToNewAccommodation(accommodation)
		response.FilteredAccommodations = append(response.FilteredAccommodations, newAccommodation)
	}
	return response, nil

}
