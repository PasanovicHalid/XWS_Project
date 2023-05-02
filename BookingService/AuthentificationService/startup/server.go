package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application"
	auth_infrastructure "github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/infrastructure/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/presentation"
	authentification_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config           *Configurations
	mux              *runtime.ServeMux
	idenitityHandler *presentation.IdentityHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.AuthentificationDBHost, config.AuthentificationDBPort)
	if err != nil {
		log.Fatal(err)
	}

	identityRepository := persistance.NewIdentityRepository(mongo)
	keyRepository := persistance.NewKeyRepository(mongo)

	keyService := auth_infrastructure.NewKeyService(keyRepository)
	jwtService := auth_infrastructure.NewJwtService(keyRepository)
	passwordService := auth_infrastructure.NewPasswordService()
	identityService := application.NewIdentityService(identityRepository, keyService, passwordService, jwtService)

	keyService.GenerateNewKeyPair()

	server.idenitityHandler = presentation.NewIdentityHandler(identityService)

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	authentification_pb.RegisterAuthenticateServiceServer(grpcServer, server.idenitityHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
