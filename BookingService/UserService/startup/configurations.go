package configurations

import "os"

type Configurations struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:       os.Getenv("USER_SERVICE_PORT"),
		UserDBHost: os.Getenv("USER_DB_HOST"),
		UserDBPort: os.Getenv("USER_DB_PORT"),
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
}
