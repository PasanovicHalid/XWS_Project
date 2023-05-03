package configurations

import (
	"fmt"
	"log"
	"net"

	user_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/presentation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config      *Configurations
	mux         *runtime.ServeMux
	userHandler *presentation.UserHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.UserDBHost, config.UserDBPort)
	if err != nil {
		log.Fatalf("Failed to connect to mongo: %v", err)
	}

	userRepository := persistance.NewUserRepository(mongo)

	userService := application.NewUserService(userRepository)

	server.userHandler = presentation.NewUserHandler(userService)

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	user_pb.RegisterUserServiceServer(grpcServer, server.userHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
