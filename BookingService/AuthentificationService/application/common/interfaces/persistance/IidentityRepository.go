package persistance

import (
	"context"
	"errors"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
)

var ErrorUsernameInUse = errors.New("Username is already in use")
var ErrorIdentityNotFound = errors.New("Identity not found")
var ErrorIdentityWithUsernameDoesntExist = errors.New("User doesn't exist")
var ErrorInvalidPassword = errors.New("Password is invalid")

type IIdentityRepository interface {
	FindIdentityByUsername(ctx *context.Context, username string) (*domain.Identity, error)
	FindIdentityById(ctx *context.Context, id string) (*domain.Identity, error)
	InsertIdentity(ctx *context.Context, identity *domain.Identity) error
	UpdateIdentity(ctx *context.Context, identity *domain.Identity) error
	DeleteIdentity(ctx *context.Context, id string, sagaTimestamp int64) (string, error)
	RollbackDeleteIdentity(ctx *context.Context, id string, sagaTimestamp int64) error
	CheckIfUsernameExists(ctx *context.Context, username string) (bool, error)
	UpdateApiKey(ctx *context.Context, id string, apiKey string) error
}
