package authentification

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
)

type KeyService struct {
	keyRepository persistance.IKeyRepository
}

func NewKeyService(keyRepository persistance.IKeyRepository) *KeyService {
	return &KeyService{
		keyRepository: keyRepository,
	}
}

func (service *KeyService) GenerateNewKeyPair() error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	keyPair, err := domain.NewKeyPair()
	if err != nil {
		return err
	}

	err = service.keyRepository.InsertNewOrReplaceOldKeyPair(&ctx, keyPair)
	if err != nil {
		return err
	}

	return nil
}

func (service *KeyService) RetrieveKeyPair(ctx *context.Context) (*domain.KeyPair, error) {
	return service.keyRepository.GetKeyPair(ctx)
}
