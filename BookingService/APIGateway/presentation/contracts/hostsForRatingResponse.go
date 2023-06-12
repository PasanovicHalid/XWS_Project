package contracts

type HostsForRatingResponse struct {
	Hosts []HostForRating `json:"hosts"`
}

type HostForRating struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Rating        float64 `json:"rating"`
	AverageRating float64 `json:"averageRating"`
	RatingId      string  `json:"ratingId"`
}
