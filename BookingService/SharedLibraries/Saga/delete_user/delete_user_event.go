package delete_user

type DeleteUserCommandType int8

const (
	DeleteUserInfo DeleteUserCommandType = iota
	RollbackUserInfo
	DeleteHostLocations
	RollbackHostLocations
	DeleteGuestPreviousReservations
	RollbackGuestPreviousReservations
	DeleteHostLocationsPreviousReservations
	RollbackHostLocationsPreviousReservations
	UnknownCommand
)

type DeleteUserCommand struct {
	Type     DeleteUserCommandType
	UserInfo DeleteUserEventInfo
}

type DeleteUserReplyType int8

const (
	DeletedUserInfo DeleteUserReplyType = iota
	UserInfoNotDeleted
	DeletedHostLocations
	HostLocationsNotDeleted
	DeletedGuestPreviousReservations
	GuestPreviousReservationsNotDeleted
	DeletedHostLocationsPreviousReservations
	HostLocationsPreviousReservationsNotDeleted
	UnknownReply
)

type DeleteUserReply struct {
	Type     DeleteUserReplyType
	UserInfo DeleteUserEventInfo
}

type DeleteUserEventInfo struct {
	UserId        string
	SagaTimestamp int64
	Role          string
}
