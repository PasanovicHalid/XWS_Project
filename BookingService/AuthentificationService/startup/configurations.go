package startup

import "os"

type Configurations struct {
	Port                     string
	AuthentificationDBPort   string
	AuthentificationDBHost   string
	NatsHost                 string
	NatsPort                 string
	NatsUser                 string
	NatsPass                 string
	DeleteUserCommandSubject string
	DeleteUserReplySubject   string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                     os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
		AuthentificationDBPort:   os.Getenv("AUTHENTIFICATION_DB_PORT"),
		AuthentificationDBHost:   os.Getenv("AUTHENTIFICATION_DB_HOST"),
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
		configurations.Port = "9101"
	}
	if configurations.AuthentificationDBPort == "" {
		configurations.AuthentificationDBPort = "27017"
	}
	if configurations.AuthentificationDBHost == "" {
		configurations.AuthentificationDBHost = "localhost"
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
