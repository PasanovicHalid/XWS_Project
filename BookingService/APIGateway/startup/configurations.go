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
	if configurations.ApiGatewayDbHost == "" {
		configurations.ApiGatewayDbHost = "localhost"
	}
	if configurations.ApiGatewayDbPort == "" {
		configurations.ApiGatewayDbPort = "27017"
	}
}
