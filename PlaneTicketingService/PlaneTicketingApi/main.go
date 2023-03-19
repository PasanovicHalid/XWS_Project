package main

import (
	"context"
	"time"
)

func main() {
	port := SetupPort()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger, storeLogger := SetupLoggers()

	db := SetupDb(timeoutContext, storeLogger, logger)
	defer db.Disconnect(timeoutContext)

	//Distribute all the connections to goroutines
	server := SetupServer(port, logger)

	//Try to shutdown gracefully
	SetupGracefullShutdown(logger, server, timeoutContext)
}
