package presentation

import (
	"context"
	"errors"
	"net/http"

	grpcservices "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/gRPCServices"
	authenticatePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type UserHandler struct {
	AuthentificationAddress string
	UserAddress             string
}

func NewUserHandler(authentificationAddress string, userAddress string) Handler {
	return &UserHandler{
		AuthentificationAddress: authentificationAddress,
		UserAddress:             userAddress,
	}
}

func (handler *UserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/user/deregister/{identityId}", handler.DeleteUser)
	if err != nil {
		panic(err)
	}
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["identityId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
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
