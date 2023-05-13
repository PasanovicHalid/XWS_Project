package configurations

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/presentation"
	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/accommodation_service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config               *Configurations
	mux                  *runtime.ServeMux
	accommodationHandler *presentation.AccommodationHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.AccommodationDBHost, config.AccommodationDBPort)
	if err != nil {
		log.Fatalf("Failed to connect to mongo: %v", err)
	}

	accomodanceRepository := persistance.NewAccommodationRepository(mongo)

	accomodanceService := application.NewAccomodationService(accomodanceRepository)

	server.accommodationHandler = presentation.NewAccomodationHandler(accomodanceService)

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reservation_pb.RegisterAccommodationServiceServer(grpcServer, server.accommodationHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
