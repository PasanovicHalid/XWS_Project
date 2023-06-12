package contracts

type AccommodationsForRatingResponse struct {
	Accommodations []AccommodationForRating `json:"accommodations"`
}

type AccommodationForRating struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	Rating        float64 `json:"rating"`
	OwnerId       string  `json:"ownerId"`
	AverageRating float64 `json:"averageRating"`
	RatingId      string  `json:"ratingId"`
}
