package configurations

import (
	"fmt"
	"log"
	"net"

	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging/nats"
	user_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/presentation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config            *Configurations
	mux               *runtime.ServeMux
	userHandler       *presentation.UserHandler
	deleteUserHandler *message_queues.DeleteUserCommandHandler
}

const (
	QueueGroup = "user_service"
)

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

	emailServiceEndpoint := fmt.Sprintf("%s:%s", server.config.EmailServiceHost, server.config.EmailServicePort)

	userService := application.NewUserService(userRepository, emailServiceEndpoint)

	deleteUserCommandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	deleteUserReplyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.deleteUserHandler = server.initDeleteUserHandler(deleteUserReplyPublisher, deleteUserCommandSubscriber, userService)

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

func (server *Server) initDeleteUserHandler(publisher saga.Publisher, subscriber saga.Subscriber, userService *application.UserService) *message_queues.DeleteUserCommandHandler {
	handler, err := message_queues.NewDeleteUserCommandHandler(userService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return handler
}
