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
	config        *Configurations
	mux           *runtime.ServeMux
	ratingHandler *presentation.RatingHandler
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.RatingDBHost, config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}

	ratingRepository := persistance.NewRatingRepository(mongo)

	notificationPublisher := server.initPublisher(server.config.NotificationSubject)

	notificationSender := message_queues.NewNotificationSender(notificationPublisher)

	ratingService := application.NewRatingService(ratingRepository, notificationSender)

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
