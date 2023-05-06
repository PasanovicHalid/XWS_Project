package main

import configurations "startup"

func main() {
	configuration := configurations.NewConfigurations()
	server := configurations.NewServer(configuration)
	server.Start()
}
