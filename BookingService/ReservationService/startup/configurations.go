package configurations

import "os"

type Configurations struct {
	Port                     string
	ReservationDBHost        string
	ReservationDBPort        string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
	NotificationSubject      string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                     os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost:        os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort:        os.Getenv("RESERVATION_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
		NotificationSubject:      os.Getenv("NOTIFICATION_SUBJECT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9104"
	}
	if configurations.ReservationDBHost == "" {
		configurations.ReservationDBHost = "localhost"
	}
	if configurations.ReservationDBPort == "" {
		configurations.ReservationDBPort = "27017"
	}
	if configurations.NatsHost == "" {
		configurations.NatsHost = "localhost"
	}
	if configurations.NatsPort == "" {
		configurations.NatsPort = "4222"
	}
	if configurations.NatsUser == "" {
		configurations.NatsUser = "xws_project"
	}
	if configurations.NatsPass == "" {
		configurations.NatsPass = "xws_project"
	}
	if configurations.DeleteUserCommandSubject == "" {
		configurations.DeleteUserCommandSubject = "delete.user.command"
	}
	if configurations.DeleteUserReplySubject == "" {
		configurations.DeleteUserReplySubject = "delete.user.reply"
	}
	if configurations.NotificationSubject == "" {
		configurations.NotificationSubject = "notification"
	}
}
