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

	switch command.Type {
	case events.DeleteGuestPreviousReservations:
		log.Println("DeleteGuestPreviousReservations")
		reply.Type = events.DeletedGuestPreviousReservations
	case events.DeleteHostLocationsPreviousReservations:
		log.Println("DeleteHostLocationsPreviousReservations")
		reply.Type = events.DeletedHostLocationsPreviousReservations
		reply.Type = events.HostLocationsPreviousReservationsNotDeleted
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		handler.replyPublisher.Publish(reply)
	}
}
