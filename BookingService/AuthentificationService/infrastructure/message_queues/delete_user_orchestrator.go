package message_queues

import (
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

func (o *DeleteUserOrchestrator) Start() error {
	event := events.DeleteUserCommand{
		Type: events.DeleteUserInfo,
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
		return events.DeleteHostLocations
	case events.UserInfoNotDeleted:
		return events.RollbackUserInfo
	case events.DeletedHostLocations:
		return events.DeleteHostLocationsPreviousReservations
	case events.HostLocationsNotDeleted:
		return events.RollbackHostLocations
	case events.GuestPreviousReservationsNotDeleted:
		return events.RollbackGuestPreviousReservations
	case events.HostLocationsPreviousReservationsNotDeleted:
		return events.RollbackHostLocationsPreviousReservations
	default:
		return events.UnknownCommand
	}
}
