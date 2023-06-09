package presentation

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application/common/interfaces/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/domain"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/notifications"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
)

type ReservationHandler struct {
	reservation_pb.UnimplementedReservationServiceServer
	reservationService *application.ReservationService
	notificationSender message_queues.INotificationSender
}

func NewReservationHandler(reservationService *application.ReservationService, notificationSender message_queues.INotificationSender) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
		notificationSender: notificationSender,
	}
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
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

	reservationID, err := generateRandomString(10)
	if err != nil {
		return nil, err
	}

	reservationStatus := domain.Pending

	fmt.Print(reservationID)

	reservation := &domain.Reservation{
		Id:                   reservationID,
		AccommodationOfferId: request.AccommodationOfferId,
		CustomerId:           request.CustomerId,
		HostId:               request.HostId,
		ReservationStatus:    reservationStatus,
		NumberOfGuests:       int(request.NumberOfGuests),
		StartDateTimeUTC:     startTime,
		EndDateTimeUTC:       endTime,
	}

	err = handler.reservationService.CreateReservation(reservation)

	if err != nil {
		return nil, err
	}

	notificationInfo := notifications.NotifyUserEventInfo{
		UserId: "bezbednost.projekat2023@gmail.com",
		Role:   "Host",
	}

	notification := notifications.NotifyUserNotification{
		Type:     0,
		UserInfo: notificationInfo,
	}
	handler.notificationSender.SendNotification(&notification)

	return &reservation_pb.CreateReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *ReservationHandler) CreateReservationAutomaticly(ctx context.Context, request *reservation_pb.CreateReservationRequest) (response *reservation_pb.CreateReservationResponse, err error) {
	layout := "2006-01-02T15:04:05Z"

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

	reservationID, err := generateRandomString(10)
	if err != nil {
		return nil, err
	}

	reservationStatus := domain.Accepted

	fmt.Print(reservationID)

	reservation := &domain.Reservation{
		Id:                   reservationID,
		AccommodationOfferId: request.AccommodationOfferId,
		CustomerId:           request.CustomerId,
		HostId:               request.HostId,
		ReservationStatus:    reservationStatus,
		NumberOfGuests:       int(request.NumberOfGuests),
		StartDateTimeUTC:     startTime,
		EndDateTimeUTC:       endTime,
	}

	err = handler.reservationService.CreateReservation(reservation)

	if err != nil {
		return nil, err
	}
	notificationInfo := notifications.NotifyUserEventInfo{
		UserId: "bezbednost.projekat2023@gmail.com",
		Role:   "Host",
	}

	notification := notifications.NotifyUserNotification{
		Type:     0,
		UserInfo: notificationInfo,
	}
	handler.notificationSender.SendNotification(&notification)

	return &reservation_pb.CreateReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
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

	otherReservations, err := handler.reservationService.GetReservationsByAccommodationOfferID(reservation.AccommodationOfferId)
	if err != nil {
		return nil, err
	}
	// Set the status of other reservations to "Pending"
	for _, otherReservation := range otherReservations {
		if otherReservation.Id != request.Id {
			otherReservation.ReservationStatus = domain.Rejected

			err = handler.reservationService.UpdateReservation(otherReservation)
			if err != nil {
				return nil, err
			}
		}
	}
	// Prepare the response
	response := &reservation_pb.AcceptReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}

	notificationInfo := notifications.NotifyUserEventInfo{
		UserId: "bezbednost.projekat2023@gmail.com",
		Role:   "Host",
	}

	notification := notifications.NotifyUserNotification{
		Type:     5,
		UserInfo: notificationInfo,
	}
	handler.notificationSender.SendNotification(&notification)

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

	notificationInfo := notifications.NotifyUserEventInfo{
		UserId: "bezbednost.projekat2023@gmail.com",
		Role:   "Host",
	}

	notification := notifications.NotifyUserNotification{
		Type:     5,
		UserInfo: notificationInfo,
	}
	handler.notificationSender.SendNotification(&notification)

	return response, nil
}

func (handler *ReservationHandler) CancelReservation(ctx context.Context, request *reservation_pb.CancelReservationRequest) (*reservation_pb.CancelReservationResponse, error) {
	// Get the reservation associated with the given reservation ID
	reservation, err := handler.reservationService.GetReservationById(request.Id)
	if err != nil {
		return nil, err
	}

	// Set the reservation's status to "Canceled"
	reservation.ReservationStatus = domain.Canceled

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

	notificationInfo := notifications.NotifyUserEventInfo{
		UserId: "bezbednost.projekat2023@gmail.com",
		Role:   "Host",
	}

	notification := notifications.NotifyUserNotification{
		Type:     1,
		UserInfo: notificationInfo,
	}
	handler.notificationSender.SendNotification(&notification)

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
		Reservations: pbPendingReservations,
	}
	return response, nil
}

func (handler *ReservationHandler) GetGuestAcceptedReservations(ctx context.Context, request *reservation_pb.GetGuestAcceptedReservationsRequest) (response *reservation_pb.GetGuestAcceptedReservationsResponse, err error) {
	// Retrieve the guest ID from the request
	guestID := request.Id

	// Get all reservations for the guest
	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}

	// Filter reservations to include only those with status ACCEPTED for the guest
	pendingReservations := []*domain.Reservation{}
	for _, reservation := range reservations {
		if reservation.CustomerId == guestID && reservation.ReservationStatus == domain.Accepted {
			pendingReservations = append(pendingReservations, reservation)
		}
	}
	pbReservations := make([]*reservation_pb.Reservation, len(pendingReservations))
	for i, res := range pendingReservations {
		pbReservations[i] = &reservation_pb.Reservation{
			Id:                   res.Id,
			AccommodationOfferId: res.AccommodationOfferId,
			CustomerId:           res.CustomerId,
			HostId:               res.HostId,
			ReservationStatus:    reservation_pb.ReservationStatus_ACCEPTED,
			NumberOfGuests:       int32(res.NumberOfGuests),
			StartDateTimeUtc:     res.StartDateTimeUTC.String(),
			EndDateTimeUtc:       res.EndDateTimeUTC.String(),
		}
	}

	// Prepare the response
	response = &reservation_pb.GetGuestAcceptedReservationsResponse{
		Reservations: pbReservations,
	}
	return response, nil
}

func (handler *ReservationHandler) CheckHostIsDistinguished(ctx context.Context, request *reservation_pb.CheckHostIsDistinguishedRequest) (response *reservation_pb.CheckHostIsDistinguishedResponse, err error) {
	hostID := request.Id

	reservations, err := handler.reservationService.GetAllReservations()
	if err != nil {
		return nil, err
	}

	acceptedReservations := []*domain.Reservation{}
	canceledReservations := []*domain.Reservation{}
	allOfHostReservations := []*domain.Reservation{}
	for _, reservation := range reservations {
		if reservation.HostId == hostID {
			allOfHostReservations = append(allOfHostReservations, reservation)
			if reservation.ReservationStatus == domain.Accepted {
				acceptedReservations = append(acceptedReservations, reservation)
			}
			if reservation.ReservationStatus == domain.Canceled {
				canceledReservations = append(canceledReservations, reservation)
			}
		}
	}

	response = &reservation_pb.CheckHostIsDistinguishedResponse{
		IsDistinguished: false,
		RequestResult: &common_pb.RequestResult{
			Code:    200,
			Message: "OK",
		},
	}

	if len(allOfHostReservations) == 0 {
		response.RequestResult = &common_pb.RequestResult{
			Code:    404,
			Message: "Host not found",
		}
		return
	}

	if len(acceptedReservations) < 5 {
		return
	}

	var durationDays int64 = 0
	for _, reservation := range acceptedReservations {
		durationDays += int64(reservation.EndDateTimeUTC.Sub(reservation.StartDateTimeUTC).Hours()) / 24
	}

	if durationDays < 50 {
		return
	}

	cancelationRate := float64(len(canceledReservations)) / float64(len(allOfHostReservations))

	if cancelationRate > 0.05 {
		return
	}

	response.IsDistinguished = true
	return
}
