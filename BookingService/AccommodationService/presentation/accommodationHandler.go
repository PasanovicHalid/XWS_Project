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

func (handler *AccommodationHandler) TempServiceMethod(ctx context.Context, message *accomodancePB.TempMessage) (*common_pb.RequestResult, error) {
	handler.accomodationService.PrintSuccess(message.String())
	return &common_pb.RequestResult{
		Code:    200,
		Message: "USPESNO",
	}, nil
}
