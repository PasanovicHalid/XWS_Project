package persistance

import (
	"context"
	"errors"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
)

var ErrorKeyNotFound = errors.New("Key not found")

type IKeyRepository interface {
	GetKeyPair(ctx *context.Context) (*domain.KeyPair, error)
	InsertNewOrReplaceOldKeyPair(ctx *context.Context, keyPair *domain.KeyPair) error
}
