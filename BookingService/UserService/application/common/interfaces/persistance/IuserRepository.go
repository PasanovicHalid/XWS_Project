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
	DeleteUser(ctx *context.Context, id string, sagaTimestamp int64) error
	RollbackDeleteUser(ctx *context.Context, id string, sagaTimestamp int64) error
	GetAllUsersByIdList(ctx *context.Context, idList []string) ([]*domain.User, error)
}
