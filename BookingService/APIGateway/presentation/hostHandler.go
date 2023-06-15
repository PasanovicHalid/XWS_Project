package presentation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/application/common/interfaces/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/infrastructure/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/contracts"
	grpcservices "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/gRPCServices"
	mw "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/startup/middlewares"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/notifications"
	ratingPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
	reservationPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type HostHandler struct {
	ReservationAddress string
	RatingAddress      string
	UserAddress        string
	NotificationSender message_queues.INotificationSender
}

func NewHostHandler(reservationAddress string, ratingAddress string, userAddress string, notificationSender message_queues.INotificationSender) Handler {
	return &HostHandler{
		ReservationAddress: reservationAddress,
		RatingAddress:      ratingAddress,
		UserAddress:        userAddress,
		NotificationSender: notificationSender,
	}
}

func (handler *HostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/user/host/distinguished", handler.DistinguishedHost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/api/rating/get-hosts-for-rating", handler.GetHostsForRating)
	if err != nil {
		panic(err)
	}
}

func (handler *HostHandler) GetHostsForRating(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jwt_claims := r.Context().Value(mw.JwtContent{}).(*authentification.SignedDetails)
	id := jwt_claims.Id

	reservations, err := handler.getGuestAcceptedReservations(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	ratings, err := handler.getRatingsOfCustomer(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	hosts := make(map[string]float64, len(ratings.Ratings))
	hostIds := make([]string, 0, len(ratings.Ratings))
	ratingMap := make(map[string]string, len(ratings.Ratings))

	for _, reservation := range reservations.Reservations {
		hostIds = append(hostIds, reservation.HostId)
	}

	for _, rating := range ratings.Ratings {
		if hosts[rating.HostId] == 0 {
			hosts[rating.HostId] = rating.Rating
			ratingMap[rating.HostId] = rating.Id
		}
	}

	hostResponse, err := handler.getHostFromIdList(hostIds)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := &contracts.HostsForRatingResponse{
		Hosts: make([]contracts.HostForRating, 0, len(hostResponse.Users)),
	}

	for _, host := range hostResponse.Users {
		averageRating, err := handler.getAverageRatingForHost(host.IdentityId)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		response.Hosts = append(response.Hosts, contracts.HostForRating{
			Id:            host.IdentityId,
			Name:          host.FirstName + " " + host.LastName,
			Rating:        hosts[host.IdentityId],
			AverageRating: averageRating.Rating,
			RatingId:      ratingMap[host.IdentityId],
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (handler *HostHandler) DistinguishedHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jwt_claims := r.Context().Value(mw.JwtContent{}).(*authentification.SignedDetails)
	id := jwt_claims.Id

	user, err := handler.getUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if user.RequestResult.Code != 200 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(user.RequestResult)
		return
	}

	distinguisehdResponse, err := handler.checkIfHostHasDistinguishedReservationQualities(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
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
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if averageRatingResponse.RequestResult.Code != 200 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(averageRatingResponse.RequestResult)
		return
	}

	isDistinguished := distinguisehdResponse.IsDistinguished && averageRatingResponse.Rating >= 4.7

	if user.User.IsDistinguished != isDistinguished {
		resp, err := handler.changeHostDistinguishedStatus(id, isDistinguished)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		if resp.RequestResult.Code != 200 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp.RequestResult)
			return
		}

		handler.NotificationSender.SendNotification(&notifications.NotifyUserNotification{
			UserInfo: notifications.NotifyUserEventInfo{
				UserId:        id,
				Distinguished: isDistinguished,
				Role:          "Host",
			},
			Type: notifications.DistinguishedChanged,
		})

	}

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

func (handler *HostHandler) getRatingsOfCustomer(id string) (*ratingPB.GetAllRatingsResponse, error) {
	ratingClient := grpcservices.NewRatingClient(handler.RatingAddress)

	return ratingClient.GetAllRatingsMadeByCustomer(context.TODO(), &ratingPB.GetAllRatingsMadeByCustomerRequest{
		Id: id,
	})
}

func (handler *HostHandler) getHostFromIdList(ids []string) (*userPB.GetAllUsersResponse, error) {
	userClient := grpcservices.NewUserClient(handler.UserAddress)

	return userClient.GetAllUsers(context.TODO(), &userPB.GetAllUsersRequest{
		Ids: ids,
	})
}

func (handler *HostHandler) getGuestAcceptedReservations(id string) (*reservationPB.GetGuestAcceptedReservationsResponse, error) {
	reservationClient := grpcservices.NewReservationClient(handler.ReservationAddress)

	return reservationClient.GetGuestAcceptedReservations(context.TODO(), &reservationPB.GetGuestAcceptedReservationsRequest{
		Id: id,
	})
}

func (handler *HostHandler) getUserById(id string) (*userPB.GetUserByIdResponse, error) {
	userClient := grpcservices.NewUserClient(handler.UserAddress)

	return userClient.GetUserById(context.TODO(), &userPB.GetUserByIdRequest{
		Id: id,
	})
}

func (handler *HostHandler) changeHostDistinguishedStatus(id string, distinguished bool) (*userPB.ChangeDistinguishedStatusResponse, error) {
	userClient := grpcservices.NewUserClient(handler.UserAddress)

	return userClient.ChangeDistinguishedStatus(context.TODO(), &userPB.ChangeDistinguishedStatusRequest{
		IdentityId:      id,
		IsDistinguished: distinguished,
	})
}
