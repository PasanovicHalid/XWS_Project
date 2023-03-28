package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"planeTicketing/contracts"
	"planeTicketing/database"
	"planeTicketing/model"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightControllerDependecies struct {
	FlightCollection *database.DatabaseCollection
	TicketCollection *database.DatabaseCollection
}

var FlightController *FlightControllerDependecies

func SetupFlightControllerRoutes(router *mux.Router) {
	createFlightRouter := router.Methods(http.MethodPost).Subrouter()
	createFlightRouter.HandleFunc("/flight/create", CreateFlight)
	createFlightRouter.Use(MiddlewareFlightDeserialization)

	deleteFlightRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteFlightRouter.HandleFunc("/flight/delete/{id}", DeleteFlight)

	getFlightRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightRouter.HandleFunc("/flight/{id}", GetFlight)

	getAllFlightsRouter := router.Methods(http.MethodGet).Subrouter()
	getAllFlightsRouter.HandleFunc("/flights/all", GetAllFlights)
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

	rw.Write([]byte("funkcija"))
	rw.WriteHeader(http.StatusOK)
	rw.WriteHeader(http.StatusCreated)
}

func DeleteFlight(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	// get the flight ID from the request URL path
	params := mux.Vars(h)
	flightID, err := primitive.ObjectIDFromHex(params["id"])
	defer cancel()
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}

	// create a filter to find the flight by ID
	filter := bson.M{"_id": flightID}
	result, err := FlightController.FlightCollection.Collection.DeleteOne(ctx, filter)
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while deleting flight", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(rw, "Flight not found", http.StatusNotFound)
		return
	}

	ticketFilter := bson.M{"flightId": flightID}
	_, err = FlightController.TicketCollection.Collection.DeleteMany(ctx, ticketFilter)
	if err != nil {
		http.Error(rw, "Something failed while deleting tickets", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Flight deleted"))
}

func GetFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var flight model.Flight
	err = FlightController.FlightCollection.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&flight)
	if err != nil {
		http.Error(w, "Flight not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flight)
}

func GetAllFlights(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := FlightController.FlightCollection.Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve flights", http.StatusInternalServerError)
		return
	}

	// var flights []model.Flight
	// if err = cursor.All(ctx, &flights); err != nil {
	// 	http.Error(w, "Failed to decode flights", http.StatusInternalServerError)
	// 	return
	// }

	var flights []contracts.FlightContract
	for cursor.Next(ctx) {
		var flight model.Flight
		if err := cursor.Decode(&flight); err != nil {
			http.Error(w, "Failed to decode flight", http.StatusInternalServerError)
			return
		}

		fc := contracts.FlightContract{
			Start:               flight.StartDateTimeUTC,
			End:                 flight.EndDateTimeUTC,
			DepartureLocation:   flight.DepartureLocation,
			DestinationLocation: flight.DestinationLocation,
			PriceOfTicket:       flight.Price,
			NumberOfTickets:     flight.AvailableTickets,
		}

		flights = append(flights, fc)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flights)
}

func MiddlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flightContract := &contracts.FlightContract{}
		err := flightContract.FromJSON(h.Body)

		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			FlightController.FlightCollection.Logger.Panic(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flightContract)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func setupFlight(flightContract *contracts.FlightContract) *model.Flight {
	flight := &model.Flight{
		StartDateTimeUTC:    flightContract.Start,
		EndDateTimeUTC:      flightContract.End,
		DepartureLocation:   flightContract.DepartureLocation,
		DestinationLocation: flightContract.DestinationLocation,
		AvailableTickets:    flightContract.NumberOfTickets,
		MaxNumberOfTickets:  flightContract.NumberOfTickets,
		Price:               flightContract.PriceOfTicket,
	}
	return flight
}
