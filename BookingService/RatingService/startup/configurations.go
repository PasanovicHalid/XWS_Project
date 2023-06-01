package startup

import "os"

type Configurations struct {
	Port         string
	RatingDBHost string
	RatingDBPort string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:         os.Getenv("RATING_SERVICE_PORT"),
		RatingDBHost: os.Getenv("RATING_DB_HOST"),
		RatingDBPort: os.Getenv("RATING_DB_PORT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9107"
	}
	if configurations.RatingDBHost == "" {
		configurations.RatingDBHost = "localhost"
	}
	if configurations.RatingDBPort == "" {
		configurations.RatingDBPort = "27017"
	}
}
