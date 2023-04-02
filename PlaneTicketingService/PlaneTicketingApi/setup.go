package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"planeTicketing/controllers"
	"planeTicketing/database"
	"planeTicketing/services"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
)

func SetupPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}
	return port
}

func SetupLoggers() (*log.Logger, *log.Logger) {
	logger := log.New(os.Stdout, "[planeTicketing-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[planeTicketing-store] ", log.LstdFlags)
	return logger, storeLogger
}

func SetupServer(port string, logger *log.Logger) http.Server {
	router := controllers.SetupRouter()
	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Println("Server listening on port", port)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	return server
}

func SetupGracefullShutdown(logger *log.Logger, server http.Server, timeoutContext context.Context) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}

func SetupSecretKey() {
	services.SECRET_KEY = os.Getenv("SECRET_KEY")
	if len(services.SECRET_KEY) == 0 {
		services.SECRET_KEY = "K6YEcMCkMzTORUaD0q2_lKDhTbvTLS7b9fxObQuj0OhR9QwPVYSBNtSZqnk7lRyrQ_fGNg_O811NCteixKxbJQ"
	}
}

func SetupControllers() {
	controllers.UserController = &controllers.UserControllerDependecies{
		UserCollection: database.OpenCollection(database.MongoInstance, "user"),
	}
	controllers.FlightController = &controllers.FlightControllerDependecies{
		FlightCollection: database.OpenCollection(database.MongoInstance, "flight"),
		TicketCollection: database.OpenCollection(database.MongoInstance, "tickets"),
	}
	controllers.TicketController = &controllers.TicketControllerDependecies{
		TicketCollection: database.OpenCollection(database.MongoInstance, "tickets"),
	}
}
