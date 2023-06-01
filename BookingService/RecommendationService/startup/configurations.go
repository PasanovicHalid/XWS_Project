package startup

import "os"

type Configurations struct {
	Port                 string
	RecommendationDBHost string
	RecommendationDBPort string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                 os.Getenv("RECOMMENDATION_SERVICE_PORT"),
		RecommendationDBHost: os.Getenv("RECOMMENDATION_DB_HOST"),
		RecommendationDBPort: os.Getenv("RECOMMENDATION_DB_PORT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9106"
	}
	if configurations.RecommendationDBHost == "" {
		configurations.RecommendationDBHost = "localhost"
	}
	if configurations.RecommendationDBPort == "" {
		configurations.RecommendationDBPort = "7474"
	}
}
