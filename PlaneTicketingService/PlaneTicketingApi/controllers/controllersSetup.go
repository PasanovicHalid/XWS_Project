package controllers

import (
	"planeTicketing/middleware"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.MiddlewareContentTypeSet)

	SetupUserControllerRoutes(router)
	SetupFlightControllerRoutes(router)

	return router
}
