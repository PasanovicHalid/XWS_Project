package controllers

import (
	"context"
	"net/http"
	"planeTicketing/contracts"
	"planeTicketing/database"
	"planeTicketing/model"
	"time"

	"github.com/gorilla/mux"
)

type FlightControllerDependecies struct {
	FlightCollection *database.DatabaseCollection
}

var FlightController *FlightControllerDependecies

func SetupFlightControllerRoutes(router *mux.Router) {

}

func CreateFlight(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	flightContract := h.Context().Value(KeyProduct{}).(*contracts.FlightContract)

	flight := setupFlight(flightContract)

	result, err := FlightController.FlightCollection.Collection.InsertOne(ctx, flight)
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while creating flight", http.StatusInternalServerError)
		return
	}

	FlightController.FlightCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	rw.WriteHeader(http.StatusCreated)
}

func setupFlight(flightContract *contracts.FlightContract) *model.Flight {
	tickets := make([]model.Ticket, flightContract.NumberOfTickets)
	for i := 0; i < flightContract.NumberOfTickets; i++ {
		tickets[i] = model.Ticket{
			Price: flightContract.PriceOfTicket,
		}
	}
	flight := &model.Flight{
		StartDateTimeUTC:    flightContract.Start,
		EndDateTimeUTC:      flightContract.End,
		DepartureLocation:   flightContract.DepartureLocation,
		DestinationLocation: flightContract.DestinationLocation,
		AvailableTickets:    tickets,
		Price:               flightContract.PriceOfTicket,
	}
	return flight
}
