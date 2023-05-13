package presentation

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/domain"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
)

type ReservationHandler struct {
	reservation_pb.UnimplementedReservationServiceServer
	reservationService *application.ReservationService
}

func NewReservationHandler(reservationService *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *reservation_pb.CreateReservationRequest) (response *reservation_pb.CreateReservationResponse, err error) {
	layout := "2006-01-02T15:04:05Z"

	fmt.Print(domain.Pending)
	fmt.Print("\n\n")
	fmt.Print(domain.Accepted)
	startTimeStr := request.StartDateTimeUtc
	if startTimeStr == "" {
		return nil, errors.New("StartDateTimeUtc is empty")
	}

	endTimeStr := request.EndDateTimeUtc
	if endTimeStr == "" {
		return nil, errors.New("EndDateTimeUtc is empty")
	}

	startTime, err := time.Parse(layout, startTimeStr)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse(layout, endTimeStr)
	if err != nil {
		return nil, err
	}
	reservation := &domain.Reservation{
		Id:                   request.Id,
		AccommodationOfferId: request.AccommodationOfferId,
		CustomerId:           request.CustomerId,
		HostId:               request.HostId,
		ReservationStatus:    domain.Pending,
		NumberOfGuests:       int(request.NumberOfGuests),
		StartDateTimeUTC:     startTime,
		EndDateTimeUTC:       endTime,
	}

	err = handler.reservationService.CreateReservation(reservation)

	if err != nil {
		return nil, err
	}

	return &reservation_pb.CreateReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *ReservationHandler) CheckGuestActiveReservations(ctx context.Context, request *reservation_pb.CheckUserActiveReservationsRequest) (response *reservation_pb.CheckUserActiveReservationsResponse, err error) {
	// Retrieve the guest ID from the request
	guestID := request.Id

	// Get all reservations for the guest
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}

	// Check if any active reservations exist for the guest
	hasActiveReservations := false
	for _, reservation := range reservations {
		if reservation.CustomerId == guestID && reservation.ReservationStatus == domain.Accepted {
			hasActiveReservations = true
			break
		}
	}

	// Prepare the response
	response = &reservation_pb.CheckUserActiveReservationsResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
		HasActiveReservations: hasActiveReservations,
	}
	return response, nil
}

func (handler *ReservationHandler) CheckHostActiveReservations(ctx context.Context, request *reservation_pb.CheckUserActiveReservationsRequest) (response *reservation_pb.CheckUserActiveReservationsResponse, err error) {
	// Retrieve the host ID from the request
	hostID := request.Id

	// Get all reservations for the host
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}

	// Check if any active reservations exist for the host
	hasActiveReservations := false
	for _, reservation := range reservations {
		if reservation.HostId == hostID && reservation.ReservationStatus == domain.Accepted {
			hasActiveReservations = true
			break
		}
	}

	// Prepare the response
	response = &reservation_pb.CheckUserActiveReservationsResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
		HasActiveReservations: hasActiveReservations,
	}
	return response, nil
}
