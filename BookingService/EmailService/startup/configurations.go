package startup

import "os"

type Configurations struct {
	Port        string
	EmailDBPort string
	EmailDBHost string
	NatsHost    string
	NatsPort    string
	NatsUser    string
	NatsPass    string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:        os.Getenv("EMAIL_SERVICE_PORT"),
		EmailDBPort: os.Getenv("EMAIL_DB_PORT"),
		EmailDBHost: os.Getenv("EMAIL_DB_HOST"),
		NatsHost:    os.Getenv("NATS_HOST"),
		NatsPort:    os.Getenv("NATS_PORT"),
		NatsUser:    os.Getenv("NATS_USER"),
		NatsPass:    os.Getenv("NATS_PASS"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9105"
	}
	if configurations.EmailDBPort == "" {
		configurations.EmailDBPort = "27017"
	}
	if configurations.EmailDBHost == "" {
		configurations.EmailDBHost = "localhost"
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
}
