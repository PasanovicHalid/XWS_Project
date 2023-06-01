package startup

import "os"

type Configurations struct {
	Port string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port: os.Getenv("EMAIL_SERVICE_PORT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9105"
	}
}
