package message_queues

import (
	"log"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application"
	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/delete_user"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
)

type DeleteUserCommandHandler struct {
	reservationService *application.ReservationService
	replyPublisher     saga.Publisher
	commandSubscriber  saga.Subscriber
}

func NewDeleteUserCommandHandler(reservationService *application.ReservationService, replyPublisher saga.Publisher, commandSubscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	handler := &DeleteUserCommandHandler{
		reservationService: reservationService,
		replyPublisher:     replyPublisher,
		commandSubscriber:  commandSubscriber,
	}

	err := handler.commandSubscriber.Subscribe(handler.handle)

	if err != nil {
		return nil, err
	}

	return handler, nil
}

func (handler *DeleteUserCommandHandler) handle(command *events.DeleteUserCommand) {

	reply := &events.DeleteUserReply{}
	reply.UserInfo = command.UserInfo

	log.Println(command.Type)

	switch command.Type {
	case events.DeleteGuestPreviousReservations:
		log.Println("DeleteGuestPreviousReservations")

		hasActiveReservation, err := handler.reservationService.CheckGuestActiveReservaton(command.UserInfo.UserId)

		if err != nil || hasActiveReservation {
			log.Println("Has active reservation or error")

			reply.Type = events.GuestPreviousReservationsNotDeleted
			break
		}

		err = handler.reservationService.DeleteReservationOfGuest(command.UserInfo.UserId, command.UserInfo.SagaTimestamp)

		if err != nil {
			log.Println("Error deleting guest previous reservations")

			reply.Type = events.GuestPreviousReservationsNotDeleted
			break
		}

		reply.Type = events.DeletedGuestPreviousReservations
	case events.DeleteHostLocationsPreviousReservations:
		log.Println("DeleteHostLocationsPreviousReservations")

		hasActiveReservation, err := handler.reservationService.CheckHostActiveReservaton(command.UserInfo.UserId)

		if err != nil || hasActiveReservation {
			log.Println("Has active reservation or error")

			reply.Type = events.HostLocationsPreviousReservationsNotDeleted
			break
		}

		err = handler.reservationService.DeleteReservationOfHost(command.UserInfo.UserId, command.UserInfo.SagaTimestamp)

		if err != nil {
			log.Println("Error deleting host locations previous reservations")

			reply.Type = events.HostLocationsPreviousReservationsNotDeleted
			break
		}

		reply.Type = events.DeletedHostLocationsPreviousReservations
	case events.RollbackGuestRatings:
		log.Println("RollbackGuestRatings")

		err := handler.reservationService.ReverseDeleteReservationOfGuest(command.UserInfo.UserId, command.UserInfo.SagaTimestamp)

		if err != nil {
			log.Println("Error rolling back guest ratings")
		}

		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		handler.replyPublisher.Publish(reply)
	}
}
