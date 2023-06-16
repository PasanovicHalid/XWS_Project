package contracts

type GetRatingsForAccommodationResponse struct {
	Ratings []*Rating `json:"ratings"`
}

type Rating struct {
	Id            string  `json:"id"`
	UserId        string  `json:"-"`
	User          string  `json:"user"`
	Rating        float64 `json:"rating"`
	TimeSubmitted string  `json:"timeSubmitted"`
}
