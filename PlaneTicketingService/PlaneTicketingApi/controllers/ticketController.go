package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"planeTicketing/contracts"
	"planeTicketing/database"
	"planeTicketing/model"
	"time"
	"fmt"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketControllerDependecies struct {
	TicketCollection *database.DatabaseCollection
	FlightCollection *database.DatabaseCollection
}



var TicketController *TicketControllerDependecies


func SetupTicketControllerRoutes(router *mux.Router) {
	purchaseTicketRouter := router.Methods(http.MethodPost).Subrouter()
	purchaseTicketRouter.HandleFunc("/purchase-ticket", PurchaseTicket)
	purchaseTicketRouter.Use(MiddlewareTicketDeserialization)

	getAllTicketsForCustomer := router.Methods(http.MethodGet).Subrouter()
    getAllTicketsForCustomer.HandleFunc("/all-tickets-for-customer/{id}", GetTicketsForCustomer)
}

func MiddlewareTicketDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		ticketPurchaseContract := &contracts.TicketPurchase{}
		err := ticketPurchaseContract.FromJSON(h.Body)

		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			TicketController.TicketCollection.Logger.Panic(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, ticketPurchaseContract)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}



func PurchaseTicket(rw http.ResponseWriter, h *http.Request) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	ticketPurchaseContract := h.Context().Value(KeyProduct{}).(*contracts.TicketPurchase)
	fmt.Printf("Processing customer %v\n", ticketPurchaseContract.CustomerId)
	fmt.Printf("Processing customer %v\n", ticketPurchaseContract.FlightId)
	ticket := setupTicket(ticketPurchaseContract)
	for i:=0; i<ticketPurchaseContract.NumberOfPurchasedTickets; i++  {
		
		result, err := TicketController.TicketCollection.Collection.InsertOne(ctx, ticket)
		defer cancel()

		if err != nil {
			http.Error(rw, "Something failed while purchasing ticket", http.StatusInternalServerError)
			return
		}
		TicketController.TicketCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)

		


	}
	defer cancel()

	var flight model.Flight
	err := FlightController.FlightCollection.Collection.FindOne(ctx, bson.M{"_id": ticket.FlightId}).Decode(&flight)
	if err != nil {
		http.Error(rw, "Flight not found", http.StatusNotFound)
		return
	}	

	newAvailableTickets := flight.AvailableTickets - ticketPurchaseContract.NumberOfPurchasedTickets

	

	_,err1 := FlightController.FlightCollection.Collection.UpdateOne(
		ctx,
		bson.M{"_id": flight.Id},
		bson.M{"$set": bson.M{"availableTickets": newAvailableTickets}},
	)
	if err1 != nil {
		http.Error(rw, "Failed to update flight", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.WriteHeader(http.StatusCreated)
	response := Response{Message: "Ticket purchased"}
	json.NewEncoder(rw).Encode(response)
}

func setupTicket(ticketPurchase *contracts.TicketPurchase) *model.Ticket {
	ticket := &model.Ticket{
		Price:    ticketPurchase.PriceOfTicket,
		FlightId: ticketPurchase.FlightId,
		CustomerId: ticketPurchase.CustomerId,
	}
	return ticket
}



func GetTicketsForCustomer(rw http.ResponseWriter, h *http.Request) {
    // get the customer ID from the request
    params := mux.Vars(h)
    customerID, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        http.Error(rw, "Invalid ID", http.StatusBadRequest)
        return
    }
    
    // create a filter to find all tickets for the customer
    filter := bson.M{"customerId": customerID}
    
    // get all tickets for the customer
    var tickets []model.Ticket
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	cursor, err := FlightController.TicketCollection.Collection.Find(ctx, filter)
	if err != nil {
		http.Error(rw, "Something went wrong while fetching tickets", http.StatusInternalServerError)
		return
	}
	err = cursor.All(ctx, &tickets)
	if err != nil {
		http.Error(rw, "Something went wrong while decoding tickets", http.StatusInternalServerError)
		return
	}
    
    // create an array of TicketInfoContract from the tickets
    var ticketInfo []contracts.TicketInfoContract
    for _, t := range tickets {
		fmt.Printf("Processing ticket %v\n", t.Id)
		fmt.Printf("Processing ticket flight id %v\n", t.FlightId)
        // find the flight associated with the ticket
        var flight model.Flight
        ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()
        err := FlightController.FlightCollection.Collection.FindOne(ctx, bson.M{"_id": t.FlightId}).Decode(&flight)
        if err != nil {
            http.Error(rw, "Something went wrong while fetching flight for ticket", http.StatusInternalServerError)
            return
        }
        
        // create a TicketInfoContract from the ticket and flight information
        info := contracts.TicketInfoContract{
            Start: flight.StartDateTimeUTC,
            End: flight.EndDateTimeUTC,
            DepartureLocation: flight.DepartureLocation,
            DestinationLocation: flight.DestinationLocation,
            PriceOfTicket: t.Price,
        }
        ticketInfo = append(ticketInfo, info)
    }
    
    // return the array of TicketInfoContract
    rw.WriteHeader(http.StatusOK)
    json.NewEncoder(rw).Encode(ticketInfo)
}