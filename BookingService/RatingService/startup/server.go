package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/presentation"
	rating_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config        *Configurations
	mux           *runtime.ServeMux
	ratingHandler *presentation.RatingHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	server.ratingHandler = presentation.NewRatingHandler()

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	rating_pb.RegisterRatingServiceServer(grpcServer, server.ratingHandler)

	log.Println("Starting gRPC server on port " + server.config.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
