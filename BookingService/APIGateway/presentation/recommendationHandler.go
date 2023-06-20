package presentation

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	contracts "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/contracts"
)

type RecommendationHandler struct {
	FlightAddress string
}

func NewRecommendationHandler(flightAddress string) *RecommendationHandler {
	return &RecommendationHandler{
		FlightAddress: flightAddress,
	}
}

func (handler *RecommendationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/temp/recommendation", handler.GetRatingForAccommodation)
	if err != nil {
		panic(err)
	}
}

func (handler *RecommendationHandler) GetRatingForAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	fmt.Println("API Endpoint: /api/temp/recommendation")

	filter := &contracts.FlightFilter{}
	if err := filter.FromJSON(r.Body); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	reader, err := filter.ToReader()
	if err != nil {
		fmt.Println(err)
	}
	// Create a new request
	req, err := http.NewRequest("POST", "http://localhost:9000/flights/filter/departures", reader)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to make request:", err)
		http.Error(w, "Failed to make request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
