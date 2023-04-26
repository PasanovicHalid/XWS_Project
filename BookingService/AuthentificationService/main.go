package main

import startup "github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/startup"

func main() {
	configuration := startup.NewConfigurations()
	server := startup.NewServer(configuration)
	server.Start()
}
