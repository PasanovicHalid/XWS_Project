package grpcservices

import (
	"log"

	authenticatePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	reservationPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(address string) userPB.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatal("Failed to start gRPC connection to User service: %v", err)
	}
	return userPB.NewUserServiceClient(conn)

}

func NewAuthenticateClient(address string) authenticatePB.AuthenticateServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatal("Failed to start gRPC connection to Authentification service: %v", err)
	}
	return authenticatePB.NewAuthenticateServiceClient(conn)

}

func NewReservationClient(address string) reservationPB.ReservationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatal("Failed to start gRPC connection to Reservation service: %v", err)
	}
	return reservationPB.NewReservationServiceClient(conn)

}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
