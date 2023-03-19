package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"planeTicketing/controllers"
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
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
