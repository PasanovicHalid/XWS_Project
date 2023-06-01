package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/EmailService/presentation"
	email_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/email_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config       *Configurations
	mux          *runtime.ServeMux
	emailHandler *presentation.EmailHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	server.emailHandler = presentation.NewEmailHandler()

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	email_pb.RegisterEmailServiceServer(grpcServer, server.emailHandler)

	log.Println("Starting gRPC server on port " + server.config.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
