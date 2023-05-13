package configurations

import (
	"fmt"
	"log"
	"net"

	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/presentation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config             *Configurations
	mux                *runtime.ServeMux
	reservationHandler *presentation.ReservationHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.ReservationDBHost, config.ReservationDBPort)
	if err != nil {
		log.Fatalf("Failed to connect to mongo: %v", err)
	}

	reservationRepository := persistance.NewReservationRepository(mongo)

	reservationService := application.NewReservationService(reservationRepository)

	server.reservationHandler = presentation.NewReservationHandler(reservationService)

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reservation_pb.RegisterReservationServiceServer(grpcServer, server.reservationHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
