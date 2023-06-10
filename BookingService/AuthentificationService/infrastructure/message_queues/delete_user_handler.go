package message_queues

import (
	"log"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application"
	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/delete_user"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
)

type DeleteUserCommandHandler struct {
	identityService   *application.IdentityService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteUserCommandHandler(identityService *application.IdentityService, replyPublisher saga.Publisher, commandSubscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	handler := &DeleteUserCommandHandler{
		identityService:   identityService,
		replyPublisher:    replyPublisher,
		commandSubscriber: commandSubscriber,
	}

	err := handler.commandSubscriber.Subscribe(handler.handle)

	if err != nil {
		return nil, err
	}

	return handler, nil
}

func (handler *DeleteUserCommandHandler) handle(command *events.DeleteUserCommand) {
	switch command.Type {
	case events.RollbackUserInfo:
		fallthrough
	case events.RollbackHostLocations:
		fallthrough
	case events.RollbackGuestPreviousReservations:
		fallthrough
	case events.RollbackHostLocationsPreviousReservations:
		log.Println("Rollback")

		err := handler.identityService.RollbackDeleteIdentity(command.UserInfo.UserId, command.UserInfo.SagaTimestamp)

		if err != nil {
			log.Println("Error deleting user info")
		}
	}
}
