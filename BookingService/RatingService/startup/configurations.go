package startup

import "os"

type Configurations struct {
	Port                     string
	RatingDBHost             string
	RatingDBPort             string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	NotificationSubject      string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                     os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost:             os.Getenv("RATING_DB_HOST"),
		RatingDBPort:             os.Getenv("RATING_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		NotificationSubject:      os.Getenv("NOTIFICATION_SUBJECT"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9107"
	}
	if configurations.RatingDBHost == "" {
		configurations.RatingDBHost = "localhost"
	}
	if configurations.RatingDBPort == "" {
		configurations.RatingDBPort = "27017"
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
	if configurations.NotificationSubject == "" {
		configurations.NotificationSubject = "notification"
	}
	if configurations.DeleteUserCommandSubject == "" {
		configurations.DeleteUserCommandSubject = "delete.user.command"
	}
	if configurations.DeleteUserReplySubject == "" {
		configurations.DeleteUserReplySubject = "delete.user.reply"
	}
}
