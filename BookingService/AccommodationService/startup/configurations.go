package configurations

import "os"

type Configurations struct {
	Port                     string
	AccommodationDBHost      string
	AccommodationDBPort      string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                     os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost:      os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort:      os.Getenv("ACCOMMODATION_DB_PORT"),
		NatsHost:                 os.Getenv("NATS_HOST"),
		NatsPort:                 os.Getenv("NATS_PORT"),
		NatsUser:                 os.Getenv("NATS_USER"),
		NatsPass:                 os.Getenv("NATS_PASS"),
		DeleteUserCommandSubject: os.Getenv("DELETE_USER_COMMAND_SUBJECT"),
		DeleteUserReplySubject:   os.Getenv("DELETE_USER_REPLY_SUBJECT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9103"
	}
	if configurations.AccommodationDBHost == "" {
		configurations.AccommodationDBHost = "localhost"
	}
	if configurations.AccommodationDBPort == "" {
		configurations.AccommodationDBPort = "27017"
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
}
