package main

import (
	"fmt"

	configurations "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/Startup"
)

func main() {
	configuration := configurations.NewConfigurations()
	server := configurations.NewServer(configuration)
	fmt.Println("Server is running on port: " + configuration.Port)
	server.Start()
}
