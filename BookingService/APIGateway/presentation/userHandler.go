package presentation

import (
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

}
