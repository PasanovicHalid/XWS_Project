package presentation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/infrastructure/authentification"
	grpcservices "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/gRPCServices"
	mw "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/startup/middlewares"
	ratingPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
	reservationPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HostHandler struct {
	ReservationAddress string
	RatingAddress      string
}

func NewHostHandler(reservationAddress string, ratingAddress string) Handler {
	return &HostHandler{
		ReservationAddress: reservationAddress,
		RatingAddress:      ratingAddress,
	}
}

func (handler *HostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/user/host/distinguished", handler.DistinguishedHost)
	if err != nil {
		panic(err)
	}
}

func (handler *HostHandler) DistinguishedHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jwt_claims := r.Context().Value(mw.JwtContent{}).(*authentification.SignedDetails)
	id := jwt_claims.Id

	distinguisehdResponse, err := handler.checkIfHostHasDistinguishedReservationQualities(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if distinguisehdResponse.RequestResult.Code != 200 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(distinguisehdResponse.RequestResult)
		return
	}

	averageRatingResponse, err := handler.getAverageRatingForHost(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if averageRatingResponse.RequestResult.Code != 200 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(averageRatingResponse.RequestResult)
		return
	}

	isDistinguished := distinguisehdResponse.IsDistinguished && averageRatingResponse.Rating >= 4.7

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"distinguished": isDistinguished}
	json.NewEncoder(w).Encode(response)
}

func (handler *HostHandler) checkIfHostHasDistinguishedReservationQualities(id string) (*reservationPB.CheckHostIsDistinguishedResponse, error) {
	reservationClient := grpcservices.NewReservationClient(handler.ReservationAddress)

	return reservationClient.CheckHostIsDistinguished(context.TODO(), &reservationPB.CheckHostIsDistinguishedRequest{
		Id: id,
	})
}

func (handler *HostHandler) getAverageRatingForHost(id string) (*ratingPB.GetAverageRatingForHostResponse, error) {
	ratingClient := grpcservices.NewRatingClient(handler.RatingAddress)

	return ratingClient.GetAverageRatingForHost(context.TODO(), &ratingPB.GetAverageRatingForHostRequest{
		Id: id,
	})
}
