package main

import (
	"context"
	"planeTicketing/database"
	"time"
)

func main() {
	port := SetupPort()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger, storeLogger := SetupLoggers()

	database.MongoInstance = database.SetupDb(timeoutContext, storeLogger, logger)
	defer database.MongoInstance.Disconnect(timeoutContext)

	database.OpenCollection(database.MongoInstance, "user")
	//Distribute all the connections to goroutines
	server := SetupServer(port, logger)

	//Try to shutdown gracefully
	SetupGracefullShutdown(logger, server, timeoutContext)
}
