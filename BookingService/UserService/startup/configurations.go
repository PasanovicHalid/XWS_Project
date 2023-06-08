package configurations

import "os"

type Configurations struct {
	Port                     string
	UserDBHost               string
	UserDBPort               string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                     os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:               os.Getenv("USER_DB_HOST"),
		UserDBPort:               os.Getenv("USER_DB_PORT"),
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
		configurations.Port = "9102"
	}
	if configurations.UserDBHost == "" {
		configurations.UserDBHost = "localhost"
	}
	if configurations.UserDBPort == "" {
		configurations.UserDBPort = "27017"
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
