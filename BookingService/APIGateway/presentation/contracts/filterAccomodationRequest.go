package contracts

import "time"

type FilterAccomodationRequest struct {
	Location            string
	MinPrice            int32
	MaxPrice            int32
	GuestNumber         int32
	From                time.Time
	To                  time.Time
	FilterByRating      bool
	RatingBottom        float64
	RatingTop           float64
	HostIsDistinguished bool
	FilterByBenefits    bool
	Wifi                bool
	Kitchen             bool
	AirConditioner      bool
	Parking             bool
}
