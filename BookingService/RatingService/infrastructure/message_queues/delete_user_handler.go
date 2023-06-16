package message_queues

import (
	"log"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application"
	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/delete_user"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
)

type DeleteUserCommandHandler struct {
	ratingService     *application.RatingService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteUserCommandHandler(ratingService *application.RatingService, replyPublisher saga.Publisher, commandSubscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	handler := &DeleteUserCommandHandler{
		ratingService:     ratingService,
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
	case events.DeleteGuestRatings:
		log.Println("DeleteGuestRatings")
		err := handler.ratingService.DeleteAllRatingsMadeByCustomer(command.UserInfo.UserId)

		if err != nil {
			log.Println("Error deleting guest ratings")

			reply.Type = events.GuestRatingsNotDeleted
			break
		}

		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		handler.replyPublisher.Publish(reply)
	}
}
