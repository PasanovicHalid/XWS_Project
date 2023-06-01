package main

import "github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/startup"

func main() {
	configuration := startup.NewConfigurations()
	server := startup.NewServer(configuration)
	server.Start()
}
