package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/infrastructure/message_queues"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/presentation"
	saga "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging"
	"github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/Saga/messaging/nats"
	rating_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config            *Configurations
	mux               *runtime.ServeMux
	ratingHandler     *presentation.RatingHandler
	deleteUserHandler *message_queues.DeleteUserCommandHandler
}

const (
	QueueGroup = "rating_service"
)

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.RatingDBHost, config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}

	deleteUserCommandSubscriber := server.initSubscriber(server.config.DeleteUserCommandSubject, QueueGroup)
	deleteUserReplyPublisher := server.initPublisher(server.config.DeleteUserReplySubject)

	ratingRepository := persistance.NewRatingRepository(mongo)

	notificationPublisher := server.initPublisher(server.config.NotificationSubject)

	notificationSender := message_queues.NewNotificationSender(notificationPublisher)

	ratingService := application.NewRatingService(ratingRepository, notificationSender)

	server.deleteUserHandler = server.initDeleteUserHandler(deleteUserReplyPublisher, deleteUserCommandSubscriber, ratingService)

	server.ratingHandler = presentation.NewRatingHandler(ratingService)

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

func (server *Server) initDeleteUserHandler(publisher saga.Publisher, subscriber saga.Subscriber, ratingService *application.RatingService) *message_queues.DeleteUserCommandHandler {
	handler, err := message_queues.NewDeleteUserCommandHandler(ratingService, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return handler
}
