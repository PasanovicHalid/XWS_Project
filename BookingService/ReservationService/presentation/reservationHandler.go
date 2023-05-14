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

func (handler *ReservationHandler) GetHostPendingReservations(ctx context.Context, request *reservation_pb.GetHostPendingReservationsRequest) (response *reservation_pb.GetHostPendingReservationsResponse, err error) {
	// Retrieve the host ID from the request
	hostID := request.Id

	// Get all reservations for the host
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}

	// Filter reservations to include only those with status PENDING for the host
	pendingReservations := []*domain.Reservation{}
	for _, reservation := range reservations {
		if reservation.HostId == hostID && reservation.ReservationStatus == domain.Pending {
			pendingReservations = append(pendingReservations, reservation)
		}
	}
	pbPendingReservations := make([]*reservation_pb.Reservation, len(pendingReservations))
	for i, res := range pendingReservations {
		pbPendingReservations[i] = &reservation_pb.Reservation{
			Id:                   res.Id,
			AccommodationOfferId: res.AccommodationOfferId,
			CustomerId:           res.CustomerId,
			HostId:               res.HostId,
			ReservationStatus:    reservation_pb.ReservationStatus_PENDING,
			NumberOfGuests:       int32(res.NumberOfGuests),
			StartDateTimeUtc:     res.StartDateTimeUTC.String(),
			EndDateTimeUtc:       res.EndDateTimeUTC.String(),
		}
	}

	// Prepare the response
	response = &reservation_pb.GetHostPendingReservationsResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
		Reservations: pbPendingReservations,
	}
	return response, nil
}

func (handler *ReservationHandler) AcceptReservation(ctx context.Context, request *reservation_pb.AcceptReservationRequest) (*reservation_pb.AcceptReservationResponse, error) {
	// Retrieve the reservation ID from the request
	reservationID := request.Id

	reservation, err := handler.reservationService.GetReservationById(reservationID)
	reservation.ReservationStatus = domain.Accepted
	handler.reservationService.UpdateReservation(reservation)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	response := &reservation_pb.AcceptReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}

	return response, nil
}

func (handler *ReservationHandler) RejectReservation(ctx context.Context, request *reservation_pb.RejectReservationRequest) (*reservation_pb.RejectReservationResponse, error) {
	// Retrieve the reservation ID from the request
	reservationID := request.Id

	reservation, err := handler.reservationService.GetReservationById(reservationID)
	reservation.ReservationStatus = domain.Rejected
	handler.reservationService.UpdateReservation(reservation)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	response := &reservation_pb.RejectReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}

	return response, nil
}

func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *reservation_pb.CancelReservationRequest) (*reservation_pb.CancelReservationResponse, error) {
	// Get the reservation associated with the given reservation ID
	reservation, err := handler.reservationService.GetReservationById(request.Id)
	if err != nil {
		return nil, err
	}

	// Set the reservation's status to "Rejected"
	reservation.ReservationStatus = domain.Rejected

	// Update the reservation
	err = handler.reservationService.UpdateReservation(reservation)
	if err != nil {
		return nil, err
	}

	// Get all other reservations with the same accommodationOfferID
	otherReservations, err := handler.reservationService.GetReservationsByAccommodationOfferID(reservation.AccommodationOfferId)
	if err != nil {
		return nil, err
	}
	// Set the status of other reservations to "Pending"
	for _, otherReservation := range otherReservations {
		if otherReservation.Id != request.Id {
			otherReservation.ReservationStatus = domain.Pending

			err = handler.reservationService.UpdateReservation(otherReservation)
			if err != nil {
				return nil, err
			}
		}
	}

	// Prepare the response
	response := &reservation_pb.CancelReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}

	return response, nil
}

func (handler *ReservationHandler) DeleteReservation(ctx context.Context, request *reservation_pb.DeleteReservationRequest) (*reservation_pb.DeleteReservationResponse, error) {
	// Retrieve the reservation ID from the request
	reservationID := request.Id

	// Make gRPC call to the reservation service to delete the reservation
	err := handler.reservationService.DeleteReservation(reservationID)
	if err != nil {
		return nil, err
	}

	// Prepare the response
	response := &reservation_pb.DeleteReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}

	return response, nil
}

func (handler *ReservationHandler) GetGuestPendingReservations(ctx context.Context, request *reservation_pb.GetGuestPendingReservationsRequest) (response *reservation_pb.GetGuestPendingReservationsResponse, err error) {
	// Retrieve the guest ID from the request
	guestID := request.Id

	// Get all reservations for the guest
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}

	// Filter reservations to include only those with status PENDING for the guest
	pendingReservations := []*domain.Reservation{}
	for _, reservation := range reservations {
		if reservation.CustomerId == guestID && reservation.ReservationStatus == domain.Pending {
			pendingReservations = append(pendingReservations, reservation)
		}
	}
	pbPendingReservations := make([]*reservation_pb.Reservation, len(pendingReservations))
	for i, res := range pendingReservations {
		pbPendingReservations[i] = &reservation_pb.Reservation{
			Id:                   res.Id,
			AccommodationOfferId: res.AccommodationOfferId,
			CustomerId:           res.CustomerId,
			HostId:               res.HostId,
			ReservationStatus:    reservation_pb.ReservationStatus_PENDING,
			NumberOfGuests:       int32(res.NumberOfGuests),
			StartDateTimeUtc:     res.StartDateTimeUTC.String(),
			EndDateTimeUtc:       res.EndDateTimeUTC.String(),
		}
	}

	// Prepare the response
	response = &reservation_pb.GetGuestPendingReservationsResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
		Reservations: pbPendingReservations,
	}
	return response, nil
}
