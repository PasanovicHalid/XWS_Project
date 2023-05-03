package persistance

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/domain"
)

type IKeyRepository interface {
	GetKey(ctx *context.Context) (*domain.Key, error)
	SaveKey(ctx *context.Context, key *domain.Key) error
}
