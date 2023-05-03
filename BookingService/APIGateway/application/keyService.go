package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/domain"
)

type KeyService struct {
	keyRepository persistance.IKeyRepository
}

func NewKeyService(keyRepository persistance.IKeyRepository) *KeyService {
	return &KeyService{
		keyRepository: keyRepository,
	}
}

func (service *KeyService) GetKey() (*domain.Key, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.keyRepository.GetKey(&ctx)
}

func (service *KeyService) SaveKey(key *domain.Key) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.keyRepository.SaveKey(&ctx, key)
}
