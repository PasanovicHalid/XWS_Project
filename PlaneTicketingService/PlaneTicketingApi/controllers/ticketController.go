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
)

type TicketControllerDependecies struct {
	TicketCollection *database.DatabaseCollection
}



var TicketController *TicketControllerDependecies

func SetupTicketControllerRoutes(router *mux.Router) {
	purchaseTicketRouter := router.Methods(http.MethodPost).Subrouter()
	purchaseTicketRouter.HandleFunc("/purchase-ticket", PurchaseTicket)
	purchaseTicketRouter.Use(MiddlewareTicketDeserialization)
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

	ticket := setupTicket(ticketPurchaseContract)
	// for i:=0; i<ticketPurchaseContract.NumberOfPurchasedTickets; i++ {
		
		result, err := TicketController.TicketCollection.Collection.InsertOne(ctx, ticket)
		defer cancel()

		if err != nil {
			http.Error(rw, "Something failed while purchasing ticket", http.StatusInternalServerError)
			return
		}
		TicketController.TicketCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	// }
	// defer cancel()
	// result, err := FlightController.FlightCollection.Collection.InsertOne(ctx, flight)
	// defer cancel()

	// if err != nil {
	// 	http.Error(rw, "Something failed while creating flight", http.StatusInternalServerError)
	// 	return
	// }

	// TicketController.TicketCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)

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