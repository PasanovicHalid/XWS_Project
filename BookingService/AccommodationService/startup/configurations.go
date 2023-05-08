package configurations

import "os"

type Configurations struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
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
}
