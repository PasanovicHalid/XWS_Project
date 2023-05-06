package configurations

import "os"

type Configurations struct {
	Port               string
	AccomodationDBHost string
	AccomodationDBPort string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:               os.Getenv("ACCOMODATION_SERVICE_PORT"),
		AccomodationDBHost: os.Getenv("ACCOMODATION_DB_HOST"),
		AccomodationDBPort: os.Getenv("ACCOMODATION_DB_PORT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9103"
	}
	if configurations.AccomodationDBHost == "" {
		configurations.AccomodationDBHost = "localhost"
	}
	if configurations.AccomodationDBPort == "" {
		configurations.AccomodationDBPort = "27017"
	}
}
