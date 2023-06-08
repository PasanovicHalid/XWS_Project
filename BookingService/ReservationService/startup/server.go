package configurations

import (
	"fmt"
	"log"
	"net"

	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging/nats"
	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/presentation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config             *Configurations
	mux                *runtime.ServeMux
	reservationHandler *presentation.ReservationHandler
	deleteUserHandler  *message_queues.DeleteUserCommandHandler
}

const (
	QueueGroup = "reservation_service"
)

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

	deleteUserCommandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	deleteUserReplyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)
	server.deleteUserHandler = server.initDeleteUserHandler(deleteUserReplyPublisher, deleteUserCommandSubscriber, reservationService)

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

func (server *Server) initDeleteUserHandler(publisher saga.Publisher, subscriber saga.Subscriber, reservationService *application.ReservationService) *message_queues.DeleteUserCommandHandler {
	handler, err := message_queues.NewDeleteUserCommandHandler(reservationService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return handler
}
