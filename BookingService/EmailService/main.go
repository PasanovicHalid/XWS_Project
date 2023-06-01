package main

import "github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/startup"

func main() {
	configuration := startup.NewConfigurations()
	server := startup.NewServer(configuration)
	server.Start()
}
