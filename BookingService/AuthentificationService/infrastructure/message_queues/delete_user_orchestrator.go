package message_queues

import (
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
	events "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/delete_user"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
)

type DeleteUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeleteUserOrchestrator(commandPublisher saga.Publisher, replySubscriber saga.Subscriber) (*DeleteUserOrchestrator, error) {
	o := &DeleteUserOrchestrator{
		commandPublisher: commandPublisher,
		replySubscriber:  replySubscriber,
	}

	err := o.replySubscriber.Subscribe(o.handle)

	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *DeleteUserOrchestrator) Start(userInfo events.DeleteUserEventInfo) error {
	event := events.DeleteUserCommand{
		Type:     events.DeleteUserInfo,
		UserInfo: userInfo,
	}
	return o.commandPublisher.Publish(event)
}

func (o *DeleteUserOrchestrator) handle(reply *events.DeleteUserReply) {
	command := events.DeleteUserCommand{
		UserInfo: reply.UserInfo,
	}
	command.Type = o.nextCommandType(reply)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteUserOrchestrator) nextCommandType(reply *events.DeleteUserReply) events.DeleteUserCommandType {
	switch reply.Type {
	case events.DeletedUserInfo:
		if reply.UserInfo.Role == domain.Role_Host {
			return events.DeleteHostLocations
		}
		if reply.UserInfo.Role == domain.Role_Guest {
			return events.DeleteGuestPreviousReservations
		}
		return events.UnknownCommand
	case events.UserInfoNotDeleted:
		return events.RollbackUserInfo
	case events.DeletedHostLocations:
		return events.DeleteHostLocationsPreviousReservations
	case events.HostLocationsNotDeleted:
		return events.RollbackHostLocations
	case events.HostLocationsPreviousReservationsNotDeleted:
		return events.RollbackHostLocationsPreviousReservations
	case events.DeletedGuestPreviousReservations:
		return events.DeleteGuestRatings
	case events.GuestPreviousReservationsNotDeleted:
		return events.RollbackGuestPreviousReservations
	case events.GuestRatingsNotDeleted:
		return events.RollbackGuestRatings
	default:
		return events.UnknownCommand
	}
}
