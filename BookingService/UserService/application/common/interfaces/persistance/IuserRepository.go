package persistance

import (
	"context"
	"errors"

	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/domain"
)

var ErrorUserNotFound = errors.New("User not found")

type IUserRepository interface {
	FindUserById(ctx *context.Context, id string) (*domain.User, error)
	CreateUser(ctx *context.Context, user *domain.User) error
	UpdateUser(ctx *context.Context, user *domain.User) error
}