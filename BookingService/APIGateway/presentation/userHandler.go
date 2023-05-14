package presentation

import (
	"context"
	"errors"
	"net/http"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/infrastructure/authentification"
	grpcservices "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/gRPCServices"
	mw "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/startup/middlewares"
	accommodation "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/accommodation_service"
	authenticatePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	reservation "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type UserHandler struct {
	AuthentificationAddress string
	UserAddress             string
	AccommodationAddress    string
	ReservationAddress      string
}

func NewUserHandler(authentificationAddress string, userAddress string, accommodationAddress string, reservationAddress string) Handler {
	return &UserHandler{
		AuthentificationAddress: authentificationAddress,
		UserAddress:             userAddress,
		AccommodationAddress:    accommodationAddress,
		ReservationAddress:      reservationAddress,
	}
}

func (handler *UserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/user/deregister", handler.DeleteUser)
	if err != nil {
		panic(err)
	}
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jwt_claims := r.Context().Value(mw.JwtContent{}).(*authentification.SignedDetails)
	id := jwt_claims.Id

	if jwt_claims.Role == "Guest" {
		hasReservation, err := handler.checkIfGuestHasReservations(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if hasReservation {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	} else if jwt_claims.Role == "Host" {
		hasReservation, err := handler.checkIfHostHasReservations(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if hasReservation {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = handler.removeHostAccomodations(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err := handler.removeIdentity(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handler.removeUserInfo(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) removeIdentity(id string) error {
	authentificationClient := grpcservices.NewAuthenticateClient(handler.AuthentificationAddress)
	removeResponse, err := authentificationClient.Remove(context.TODO(), &authenticatePB.RemoveRequest{IdentityId: id})
	if err != nil {
		return err
	}

	if removeResponse.RequestResult.Code != 200 {
		return errors.New(removeResponse.RequestResult.Message)
	}

	return nil
}

func (handler *UserHandler) removeUserInfo(id string) error {
	userClient := grpcservices.NewUserClient(handler.UserAddress)
	deleteResponse, err := userClient.DeleteUser(context.TODO(), &userPB.DeleteUserRequest{
		IdentityId: id,
	})

	if err != nil {
		return err
	}

	if deleteResponse.RequestResult.Code != 200 {
		return errors.New(deleteResponse.RequestResult.Message)
	}

	return nil
}

func (handler *UserHandler) checkIfGuestHasReservations(id string) (bool, error) {
	reservationClient := grpcservices.NewReservationClient(handler.ReservationAddress)
	reservationsResponse, err := reservationClient.CheckGuestActiveReservations(context.TODO(), &reservation.CheckUserActiveReservationsRequest{
		Id: id,
	})

	if err != nil {
		return false, err
	}

	if reservationsResponse.RequestResult.Code != 200 {
		return false, errors.New(reservationsResponse.RequestResult.Message)
	}

	return reservationsResponse.HasActiveReservations, nil
}

func (handler *UserHandler) checkIfHostHasReservations(id string) (bool, error) {
	reservationClient := grpcservices.NewReservationClient(handler.ReservationAddress)
	reservationsResponse, err := reservationClient.CheckHostActiveReservations(context.TODO(), &reservation.CheckUserActiveReservationsRequest{
		Id: id,
	})

	if err != nil {
		return false, err
	}

	if reservationsResponse.RequestResult.Code != 200 {
		return false, errors.New(reservationsResponse.RequestResult.Message)
	}

	return reservationsResponse.HasActiveReservations, nil
}

func (handler *UserHandler) removeHostAccomodations(id string) error {
	accomodationClient := grpcservices.NewAccomodationClient(handler.AccommodationAddress)
	deleteResponse, err := accomodationClient.DeleteAllAccommodationsByOwner(context.TODO(), &accommodation.IdentityIdRequest{
		Id: id,
	})

	if err != nil {
		return err
	}

	if deleteResponse.Code != 200 {
		return errors.New(deleteResponse.Message)
	}

	return nil
}
