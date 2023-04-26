package authentification

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
)

type IKeyService interface {
	GenerateNewKeyPair() error
	RetrieveKeyPair(ctx *context.Context) (*domain.KeyPair, error)
}
