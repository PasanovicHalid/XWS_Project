package configurations

import "os"

type Configurations struct {
	Port                 string
	ApiGatewayDbHost     string
	ApiGatewayDbPort     string
	AuthentificationHost string
	AuthentificationPort string
	UserHost             string
	UserPort             string
	ReservationHost      string
	ReservationPort      string
	AccommodationHost    string
	AccommodationPort    string
	RatingHost           string
	RatingPort           string
	RecommendationHost   string
	RecommendationPort   string
	EmailHost            string
	EmailPort            string
}

func NewConfigurations() *Configurations {
	configurations := &Configurations{
		Port:                 os.Getenv("GATEWAY_PORT"),
		ApiGatewayDbHost:     os.Getenv("API_GATEWAY_DB_HOST"),
		ApiGatewayDbPort:     os.Getenv("API_GATEWAY_DB_PORT"),
		AuthentificationHost: os.Getenv("AUTHENTIFICATION_SERVICE_HOST"),
		AuthentificationPort: os.Getenv("AUTHENTIFICATION_SERVICE_PORT"),
		UserHost:             os.Getenv("USER_SERVICE_HOST"),
		UserPort:             os.Getenv("USER_SERVICE_PORT"),
		ReservationHost:      os.Getenv("RESERVATION_SERVICE_HOST"),
		ReservationPort:      os.Getenv("RESERVATION_SERVICE_PORT"),
		AccommodationHost:    os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:    os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		RatingHost:           os.Getenv("RATING_SERVICE_HOST"),
		RatingPort:           os.Getenv("RATING_SERVICE_PORT"),
		RecommendationHost:   os.Getenv("RECOMMENDATION_SERVICE_HOST"),
		RecommendationPort:   os.Getenv("RECOMMENDATION_SERVICE_PORT"),
		EmailHost:            os.Getenv("EMAIL_SERVICE_HOST"),
		EmailPort:            os.Getenv("EMAIL_SERVICE_PORT"),
	}

	configurations.initializeEnvironmentVariables()

	return configurations
}

func (configurations *Configurations) initializeEnvironmentVariables() {
	if configurations.Port == "" {
		configurations.Port = "9100"
	}
	if configurations.AuthentificationHost == "" {
		configurations.AuthentificationHost = "localhost"
	}
	if configurations.AuthentificationPort == "" {
		configurations.AuthentificationPort = "9101"
	}
	if configurations.UserHost == "" {
		configurations.UserHost = "localhost"
	}
	if configurations.UserPort == "" {
		configurations.UserPort = "9102"
	}
	if configurations.AccommodationHost == "" {
		configurations.AccommodationHost = "localhost"
	}
	if configurations.AccommodationPort == "" {
		configurations.AccommodationPort = "9103"
	}
	if configurations.ReservationHost == "" {
		configurations.ReservationHost = "localhost"
	}
	if configurations.ReservationPort == "" {
		configurations.ReservationPort = "9104"
	}
	if configurations.EmailHost == "" {
		configurations.EmailHost = "localhost"
	}
	if configurations.EmailPort == "" {
		configurations.EmailPort = "9105"
	}
	if configurations.ApiGatewayDbHost == "" {
		configurations.ApiGatewayDbHost = "localhost"
	}
	if configurations.ApiGatewayDbPort == "" {
		configurations.ApiGatewayDbPort = "27017"
	}
	if configurations.RecommendationHost == "" {
		configurations.RecommendationHost = "localhost"
	}
	if configurations.RecommendationPort == "" {
		configurations.RecommendationPort = "9106"
	}
	if configurations.RatingHost == "" {
		configurations.RatingHost = "localhost"
	}
	if configurations.RatingPort == "" {
		configurations.RatingPort = "9107"
	}
}
