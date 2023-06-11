package presentation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/contracts"
	grpcservices "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/gRPCServices"
	mw "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/startup/middlewares"
	accommodationPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/accommodation_service"
	ratingPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
	reservationPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccommodationHandler struct {
	AccommodationAddress string
	ReservationAddress   string
	RatingAddress        string
}

func NewAccommodationHandler(accommodationAddress string, reservationAddress string, ratingAddress string) Handler {
	return &AccommodationHandler{
		AccommodationAddress: accommodationAddress,
		ReservationAddress:   reservationAddress,
		RatingAddress:        ratingAddress,
	}
}

func (handler *AccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/accomodation/get-filtered-accommodations", handler.FilterAccomodations)
	if err != nil {
		panic(err)
	}
}

func (handler *AccommodationHandler) FilterAccomodations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	request := r.Context().Value(mw.RequestDecoded{}).(*contracts.FilterAccomodationRequest)

	accommodationResponse, err := handler.getAccommodationsForInitialFilter(&accommodationPB.AccommodationSearch{
		Location:         request.Location,
		StartDateTimeUtc: timestamppb.New(request.From),
		EndDateTimeUtc:   timestamppb.New(request.To),
		GuestNumber:      request.GuestNumber,
		MinPrice:         request.MinPrice,
		MaxPrice:         request.MaxPrice,
		FilterByBenefits: request.FilterByBenefits,
		Wifi:             request.Wifi,
		Parking:          request.Parking,
		Kitchen:          request.Kitchen,
		AirConditioner:   request.AirConditioner,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	filteredAccommodations := accommodationResponse.FilteredAccommodations

	filteredAccommodations, shouldReturn := handler.filterByRating(request, filteredAccommodations, w)
	if shouldReturn {
		return
	}

	if request.HostIsDistinguished {
		handler.filterByDistinguishedHost(filteredAccommodations, w)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filteredAccommodations)
		return
	}

}

func (handler *AccommodationHandler) filterByDistinguishedHost(filteredAccommodations []*accommodationPB.NewAccomodation, w http.ResponseWriter) {
	checkedHosts := map[string]bool{}

	accomodations := make([]*accommodationPB.NewAccomodation, len(filteredAccommodations))

	for _, accommodation := range filteredAccommodations {
		if val, exists := checkedHosts[accommodation.OwnerId]; exists {
			if val {
				accomodations = append(accomodations, accommodation)
			}
		} else {
			isDistinguished, shouldReturn := handler.checkIfHostIsDistinguished(accommodation, w, checkedHosts)

			if shouldReturn {
				return
			}

			if isDistinguished {
				accomodations = append(accomodations, accommodation)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accomodations)
}

func (handler *AccommodationHandler) filterByRating(request *contracts.FilterAccomodationRequest, filteredAccommodations []*accommodationPB.NewAccomodation, w http.ResponseWriter) ([]*accommodationPB.NewAccomodation, bool) {
	if request.FilterByRating {
		filteredByRating := make([]*accommodationPB.NewAccomodation, 5)

		for _, accommodation := range filteredAccommodations {
			rating, err := handler.getRatingForAccommodation(accommodation.Id)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err.Error())
				return nil, true
			}

			if rating.Rating >= request.RatingBottom && rating.Rating <= request.RatingTop {
				filteredByRating = append(filteredByRating, accommodation)
			}
		}

		filteredAccommodations = filteredByRating
	}
	return filteredAccommodations, false
}

func (handler *AccommodationHandler) checkIfHostIsDistinguished(accommodation *accommodationPB.NewAccomodation, w http.ResponseWriter, checkedHosts map[string]bool) (isDistinguished bool, shouldReturn bool) {
	hostHasInitialQualities, err := handler.checkIfHostHasDistinguishedReservationQualities(accommodation.OwnerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		isDistinguished = false
		shouldReturn = true
		return
	}

	averageRatingResponse, err := handler.getAverageRatingForHost(accommodation.OwnerId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		isDistinguished = false
		shouldReturn = true
		return
	}

	isDistinguished = hostHasInitialQualities.IsDistinguished && averageRatingResponse.Rating >= 4.7
	shouldReturn = false

	if isDistinguished {
		checkedHosts[accommodation.OwnerId] = true
	} else {
		checkedHosts[accommodation.OwnerId] = false
	}
	return
}

func (handler *AccommodationHandler) checkIfHostHasDistinguishedReservationQualities(id string) (*reservationPB.CheckHostIsDistinguishedResponse, error) {
	reservationClient := grpcservices.NewReservationClient(handler.ReservationAddress)

	return reservationClient.CheckHostIsDistinguished(context.TODO(), &reservationPB.CheckHostIsDistinguishedRequest{
		Id: id,
	})
}

func (handler *AccommodationHandler) getAverageRatingForHost(id string) (*ratingPB.GetAverageRatingForHostResponse, error) {
	ratingClient := grpcservices.NewRatingClient(handler.RatingAddress)

	return ratingClient.GetAverageRatingForHost(context.TODO(), &ratingPB.GetAverageRatingForHostRequest{
		Id: id,
	})
}

func (handler *AccommodationHandler) getRatingForAccommodation(id string) (*ratingPB.GetRatingForAccommodationResponse, error) {
	ratingClient := grpcservices.NewRatingClient(handler.RatingAddress)

	return ratingClient.GetRatingForAccommodation(context.TODO(), &ratingPB.GetRatingForAccommodationRequest{
		Id: id,
	})
}

func (handler *AccommodationHandler) getAccommodationsForInitialFilter(filter *accommodationPB.AccommodationSearch) (*accommodationPB.GetFilteredAccommodationsResponse, error) {
	accommodationClient := grpcservices.NewAccommodationClient(handler.AccommodationAddress)

	return accommodationClient.FilterAccommodations(context.TODO(), filter)
}
