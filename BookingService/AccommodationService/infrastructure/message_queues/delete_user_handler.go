package message_queues

import (
	"log"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application"
	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/delete_user"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
)

type DeleteUserCommandHandler struct {
	accommodationService *application.AccommodationService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
}

func NewDeleteUserCommandHandler(accommodationService *application.AccommodationService, replyPublisher saga.Publisher, commandSubscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	handler := &DeleteUserCommandHandler{
		accommodationService: accommodationService,
		replyPublisher:       replyPublisher,
		commandSubscriber:    commandSubscriber,
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
	case events.DeleteHostLocations:
		log.Println("DeleteHostLocations")
		reply.Type = events.DeletedHostLocations
	case events.RollbackHostLocationsPreviousReservations:
		log.Println("Rollback")
		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		handler.replyPublisher.Publish(reply)
	}
}
