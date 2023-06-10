package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/domain"
)

type UserService struct {
	userRepository persistance.IUserRepository
}

func NewUserService(userRepository persistance.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
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
