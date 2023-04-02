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

type Response struct {
	Message string `json:"message"`
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

	getFilteredFlightsRouter := router.Methods(http.MethodPost).Subrouter()
	getFilteredFlightsRouter.HandleFunc("/flights/filter", getFilteredFlights)

	getFlightCities := router.Methods(http.MethodGet).Subrouter()
	getFlightCities.HandleFunc("/flights/cities", getAllCities)
}

func CreateFlight(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	flightContract := h.Context().Value(KeyProduct{}).(*contracts.FlightContract)

	flight := setupFlightModel(flightContract)

	result, err := FlightController.FlightCollection.Collection.InsertOne(ctx, flight)
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while creating flight", http.StatusInternalServerError)
		return
	}

	FlightController.FlightCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)

	rw.WriteHeader(http.StatusOK)
	rw.WriteHeader(http.StatusCreated)
	response := Response{Message: "Flight created"}
	json.NewEncoder(rw).Encode(response)
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
	response := Response{Message: "Flight deleted"}
	json.NewEncoder(rw).Encode(response)
}

func getFlights() ([]contracts.FlightContract, error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := FlightController.FlightCollection.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err, nil
	}

	var flights []contracts.FlightContract
	for cursor.Next(ctx) {
		var flight model.Flight
		if err := cursor.Decode(&flight); err != nil {
			return nil, nil, err
		}

		fc := setupFlightContract(&flight)

		flights = append(flights, fc)
	}

	return flights, nil, nil

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
	flights, errRetrieve, errDecode := getFlights()
	if errRetrieve != nil {
		http.Error(w, "Failed to retrieve flights", http.StatusInternalServerError)
		return
	}

	if errDecode != nil {
		http.Error(w, "Failed to decode flight", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flights)
}

func getCities() []contracts.City {
	flights, _, _ := getFlights()
	var cities []contracts.City
	for _, f := range flights {
		if !contains(cities, f.DepartureLocation) {
			tempCity := contracts.City{
				Name: f.DepartureLocation,
			}
			cities = append(cities, tempCity)
		}
		if !contains(cities, f.DestinationLocation) {
			tempCity := contracts.City{
				Name: f.DestinationLocation,
			}
			cities = append(cities, tempCity)
		}
	}

	return cities
}

func getAllCities(w http.ResponseWriter, r *http.Request) {
	var cities []contracts.City
	_, errRetrieve, errDecode := getFlights()
	if errRetrieve != nil {
		http.Error(w, "Failed to retrieve flights", http.StatusInternalServerError)
		return
	}

	if errDecode != nil {
		http.Error(w, "Failed to decode flight", http.StatusInternalServerError)
		return
	}

	cities = getCities()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cities)
}

func contains(s []contracts.City, str string) bool {
	for _, v := range s {
		if v.Name == str {
			return true
		}
	}
	return false
}

func getFilteredFlights(w http.ResponseWriter, r *http.Request) {

	flightFilter := &contracts.FlightFilter{}
	err := flightFilter.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
		FlightController.FlightCollection.Logger.Panic(err)
		return
	}

	if !validateFlightFilter(flightFilter) {
		http.Error(w, "Flight filter is not valid", http.StatusBadRequest)
		return
	}

	flights, errRetrieve, errDecode := getFlights()
	if errRetrieve != nil {
		http.Error(w, "Failed to retrieve flights", http.StatusInternalServerError)
		return
	}

	if errDecode != nil {
		http.Error(w, "Failed to decode flight", http.StatusInternalServerError)
		return
	}

	var filteredFlights []contracts.FlightContract

	yourDate, _ := time.Parse("2006-01-02", flightFilter.Date)

	for _, f := range flights {
		if f.DepartureLocation == flightFilter.DepartureLocation && f.DestinationLocation == flightFilter.DestinationLocation && DateEqual(f.Start, yourDate) && f.AvailableNumberOfTickets >= flightFilter.NumberOfTickets {
			filteredFlights = append(filteredFlights, f)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filteredFlights)
}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func validateFlightFilter(filterFlight *contracts.FlightFilter) bool {
	valid := true
	date, _ := time.Parse("2006-01-02", filterFlight.Date)
	if filterFlight.NumberOfTickets > 5 || filterFlight.NumberOfTickets < 1 {
		valid = false
	}
	if date.Before(time.Now()) {
		valid = false
	}
	if !contains(getCities(), filterFlight.DepartureLocation) || !contains(getCities(), filterFlight.DestinationLocation) {
		valid = false
	}
	return valid
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

func setupFlightModel(flightContract *contracts.FlightContract) model.Flight {
	flight := model.Flight{
		StartDateTimeUTC:    flightContract.Start,
		EndDateTimeUTC:      flightContract.End,
		DepartureLocation:   flightContract.DepartureLocation,
		DestinationLocation: flightContract.DestinationLocation,
		AvailableTickets:    flightContract.AvailableNumberOfTickets,
		MaxNumberOfTickets:  flightContract.MaxNumberOfTickets,
		Price:               flightContract.PriceOfTicket,
	}
	return flight
}

func setupFlightContract(flight *model.Flight) contracts.FlightContract {
	flightContract := contracts.FlightContract{
		Id:                       flight.Id,
		Start:                    flight.StartDateTimeUTC,
		End:                      flight.EndDateTimeUTC,
		DepartureLocation:        flight.DepartureLocation,
		DestinationLocation:      flight.DestinationLocation,
		PriceOfTicket:            flight.Price,
		MaxNumberOfTickets:       flight.MaxNumberOfTickets,
		AvailableNumberOfTickets: flight.AvailableTickets,
	}
	return flightContract
}
