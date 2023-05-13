package persistance

import "go.mongodb.org/mongo-driver/mongo"

type AccommodationRepository struct {
	identities *mongo.Collection
}

const (
	DATABASE                 = "AccommodationDB"
	ACCOMMODATION_COLLECTION = "Accommodations"
)

func NewIdentityRepository(client *mongo.Client) *AccommodationRepository {
	return &AccommodationRepository{
		identities: client.Database(DATABASE).Collection(ACCOMMODATION_COLLECTION),
	}
}
