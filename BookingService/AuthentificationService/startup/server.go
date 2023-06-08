package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application"
	auth_infrastructure "github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/infrastructure/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/presentation"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging/nats"
	authentification_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config                 *Configurations
	mux                    *runtime.ServeMux
	idenitityHandler       *presentation.IdentityHandler
	deleteUserOrchestrator *message_queues.DeleteUserOrchestrator
	deleteUserHandler      *message_queues.DeleteUserCommandHandler
}

const (
	QueueGroup = "authentification_service"
)

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.AuthentificationDBHost, config.AuthentificationDBPort)
	if err != nil {
		log.Fatal(err)
	}

	deleteUserCommandPublisher := server.initPublisher(server.config.DeleteUserCommandSubject)
	deleteUserReplySubscriber := server.initSubscriber(server.config.DeleteUserReplySubject, QueueGroup)
	server.deleteUserOrchestrator = server.initDeleteUserOrchestrator(deleteUserCommandPublisher, deleteUserReplySubscriber)

	identityRepository := persistance.NewIdentityRepository(mongo)
	keyRepository := persistance.NewKeyRepository(mongo)

	keyService := auth_infrastructure.NewKeyService(keyRepository)
	jwtService := auth_infrastructure.NewJwtService(keyRepository)
	passwordService := auth_infrastructure.NewPasswordService()
	identityService := application.NewIdentityService(identityRepository, keyService, passwordService, jwtService)

	deleteUserCommandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	deleteUserReplyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.deleteUserHandler = server.initDeleteUserHandler(deleteUserReplyPublisher, deleteUserCommandSubscriber, identityService)

	keyService.GenerateNewKeyPair()

	server.idenitityHandler = presentation.NewIdentityHandler(identityService, server.deleteUserOrchestrator)

	return server
}

func (server *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	authentification_pb.RegisterAuthenticateServiceServer(grpcServer, server.idenitityHandler)

	log.Println("Starting gRPC server on port " + server.config.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initDeleteUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *message_queues.DeleteUserOrchestrator {
	orchestrator, err := message_queues.NewDeleteUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initDeleteUserHandler(publisher saga.Publisher, subscriber saga.Subscriber, identityService *application.IdentityService) *message_queues.DeleteUserCommandHandler {
	handler, err := message_queues.NewDeleteUserCommandHandler(identityService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return handler
}
