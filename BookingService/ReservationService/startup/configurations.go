package configurations

import "os"

type Configurations struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:              os.Getenv("RESERVATION_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("RESERVATION_DB_HOST"),
		ReservationDBPort: os.Getenv("RESERVATION_DB_PORT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9104"
	}
	if configurations.ReservationDBHost == "" {
		configurations.ReservationDBHost = "localhost"
	}
	if configurations.ReservationDBPort == "" {
		configurations.ReservationDBPort = "27017"
	}
}
