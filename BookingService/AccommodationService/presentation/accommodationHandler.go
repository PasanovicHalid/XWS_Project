package presentation

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application"
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
