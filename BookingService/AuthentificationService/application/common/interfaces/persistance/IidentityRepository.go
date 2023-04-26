package persistance

import (
	"context"
	"errors"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
)

var ErrorEmailInUse = errors.New("Email is already in use")
var ErrorIdentityNotFound = errors.New("Identity not found")
var ErrorIdentityWithEmailDoesntExist = errors.New("User doesn't exist")
var ErrorInvalidPassword = errors.New("Password is invalid")

type IIdentityRepository interface {
	FindIdentityByEmail(ctx *context.Context, email string) (*domain.Identity, error)
	FindIdentityById(ctx *context.Context, id string) (*domain.Identity, error)
	InsertIdentity(ctx *context.Context, identity *domain.Identity) error
	UpdateIdentity(ctx *context.Context, identity *domain.Identity) error
	DeleteIdentity(ctx *context.Context, id string) error
	CheckIfEmailExists(ctx *context.Context, email string) (bool, error)
}
