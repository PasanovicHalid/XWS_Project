package application

import (
	"context"
	"fmt"
	"log"
	"time"

	email_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/email_service"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	userRepository      persistance.IUserRepository
	emailServiceAddress string
}

func NewUserService(userRepository persistance.IUserRepository, emailServiceAddress string) *UserService {
	return &UserService{
		userRepository:      userRepository,
		emailServiceAddress: emailServiceAddress,
	}
}

func (service *UserService) GetUserById(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.userRepository.FindUserById(&ctx, id)
}

func (service *UserService) CreateUser(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	notification := &email_pb.UpdateWantedNotificationsRequest{
		Id:                       user.Email,
		CreatedRequest:           true,
		CanceledReservation:      true,
		HostRatingGiven:          true,
		AccommodationRatingGiven: true,
		ProminentHost:            true,
		HostResponded:            true,
	}
	emailService := service.initEmailServiceClient()
	fmt.Print("AAAA\n\n")
	_, err := emailService.SetWantedNotifications(ctx, notification)
	fmt.Print(err)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Print("SSSSSSSS\n\n")
	return service.userRepository.CreateUser(&ctx, user)
}

func (service *UserService) UpdateUser(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.userRepository.UpdateUser(&ctx, user)
}

func (service *UserService) DeleteUser(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.userRepository.DeleteUser(&ctx, id, sagaTimestamp)
}

func (service *UserService) RollbackDeleteUser(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.userRepository.RollbackDeleteUser(&ctx, id, sagaTimestamp)
}

func (service *UserService) GetAllUsersByIdList(idList []string) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.userRepository.GetAllUsersByIdList(&ctx, idList)
}

func (service *UserService) ChangeDistinguishedStatus(id string, status bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.userRepository.ChangeDistinguishedStatus(&ctx, id, status)
}

func (service *UserService) initEmailServiceClient() email_pb.EmailServiceClient {
	conn, err := getConnection(service.emailServiceAddress)
	if err != nil {
		log.Fatal("Failed to start gRPC connection to Email service: %v", err)
	}
	return email_pb.NewEmailServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
