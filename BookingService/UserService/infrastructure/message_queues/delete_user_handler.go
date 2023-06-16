package message_queues

import (
	"log"

	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/delete_user"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application"
)

type DeleteUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteUserCommandHandler(userService *application.UserService, replyPublisher saga.Publisher, commandSubscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	handler := &DeleteUserCommandHandler{
		userService:       userService,
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

	reply := &events.DeleteUserReply{}
	reply.UserInfo = command.UserInfo

	log.Println(command.Type)

	switch command.Type {
	case events.DeleteUserInfo:
		log.Println("DeleteUserInfo")

		err := handler.userService.DeleteUser(command.UserInfo.UserId, command.UserInfo.SagaTimestamp)

		if err != nil {
			log.Println("Error deleting user info")

			reply.Type = events.UserInfoNotDeleted
			break
		}

		reply.Type = events.DeletedUserInfo
	case events.RollbackHostLocations:
		fallthrough
	case events.RollbackGuestPreviousReservations:
		fallthrough
	case events.RollbackGuestRatings:
		fallthrough
	case events.RollbackHostLocationsPreviousReservations:
		log.Println("Rollback")

		err := handler.userService.RollbackDeleteUser(command.UserInfo.UserId, command.UserInfo.SagaTimestamp)

		if err != nil {
			log.Println("Error rolling back delete user")
		}

		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		handler.replyPublisher.Publish(reply)
	}
}
