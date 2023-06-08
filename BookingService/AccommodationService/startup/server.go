package configurations

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AccommodationService/presentation"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging/nats"
	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/accommodation_service"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config               *Configurations
	mux                  *runtime.ServeMux
	accommodationHandler *presentation.AccommodationHandler
	deleteUserHandler    *message_queues.DeleteUserCommandHandler
}

const (
	QueueGroup = "accommodation_service"
)

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

	deleteUserCommandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	deleteUserReplyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.deleteUserHandler = server.initDeleteUserHandler(deleteUserReplyPublisher, deleteUserCommandSubscriber, accomodanceService)

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

func (server *Server) initDeleteUserHandler(publisher saga.Publisher, subscriber saga.Subscriber, accommodationService *application.AccommodationService) *message_queues.DeleteUserCommandHandler {
	handler, err := message_queues.NewDeleteUserCommandHandler(accommodationService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return handler
}
