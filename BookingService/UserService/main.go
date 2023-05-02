package main

import configurations "github.com/PasanovicHalid/XWS_Project/BookingService/UserService/startup"

func main() {
	configuration := configurations.NewConfigurations()
	server := configurations.NewServer(configuration)
	server.Start()
}
