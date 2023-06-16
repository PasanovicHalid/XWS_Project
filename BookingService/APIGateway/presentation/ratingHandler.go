package presentation

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/contracts"
	grpcservices "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/gRPCServices"
	ratingPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type RatingHandler struct {
	RatingAddress string
	UserAddress   string
}

func NewRatingHandler(ratingAddress string, userAddress string) Handler {
	return &RatingHandler{
		RatingAddress: ratingAddress,
		UserAddress:   userAddress,
	}
}

func (handler *RatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/rating/get-ratings-for-accommodation/{id}", handler.GetRatingForAccommodation)
	if err != nil {
		panic(err)
	}
}

func (handler *RatingHandler) GetRatingForAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]

	ratings, err := handler.getAccommodationRatings(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := contracts.GetRatingsForAccommodationResponse{
		Ratings: make([]*contracts.Rating, 0, len(ratings.Ratings)),
	}

	if len(ratings.Ratings) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	mapRatingsUsers := make(map[string]*contracts.Rating)
	userList := make([]string, 0, len(ratings.Ratings))
	for _, rating := range ratings.Ratings {

		ratingResponse := &contracts.Rating{
			Id:            rating.Id,
			Rating:        rating.Rating,
			TimeSubmitted: rating.TimeIssued.AsTime().Format(time.RFC3339),
			UserId:        rating.UserId,
		}

		response.Ratings = append(response.Ratings, ratingResponse)

		if _, ok := mapRatingsUsers[rating.UserId]; !ok {
			mapRatingsUsers[rating.UserId] = ratingResponse
			userList = append(userList, rating.UserId)
		}
	}

	users, err := handler.getUsersByListId(userList)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if users.RequestResult.Code != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(users.RequestResult.Message)
		return
	}

	for _, user := range users.Users {
		if val, ok := mapRatingsUsers[user.IdentityId]; ok {
			val.User = user.FirstName + " " + user.LastName
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (handler *RatingHandler) getAccommodationRatings(id string) (*ratingPB.GetAllRatingsResponse, error) {
	ratingClient := grpcservices.NewRatingClient(handler.RatingAddress)

	return ratingClient.GetAllRatingsForAccommodation(context.TODO(), &ratingPB.GetAllRatingsForAccommodationRequest{
		Id: id,
	})
}

func (handler *RatingHandler) getUsersByListId(ids []string) (*userPB.GetAllUsersResponse, error) {
	userClient := grpcservices.NewUserClient(handler.UserAddress)

	return userClient.GetAllUsers(context.TODO(), &userPB.GetAllUsersRequest{
		Ids: ids,
	})
}
