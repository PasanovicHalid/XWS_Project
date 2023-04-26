package startup

import "os"

type Configurations struct {
	Port                   string
	AuthentificationDBPort string
	AuthentificationDBHost string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                   os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
		AuthentificationDBPort: os.Getenv("AUTHENTIFICATION_DB_PORT"),
		AuthentificationDBHost: os.Getenv("AUTHENTIFICATION_DB_HOST"),
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
}
