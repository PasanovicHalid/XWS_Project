package presentation

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/RatingService/domain"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	rating_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
)

type RatingHandler struct {
	rating_pb.UnimplementedRatingServiceServer
	ratingService *application.RatingService
}

func NewRatingHandler(ratingService *application.RatingService) *RatingHandler {
	return &RatingHandler{ratingService: ratingService}
}

func (h *RatingHandler) GetAllRatingsMadeByCustomer(ctx context.Context, request *rating_pb.GetAllRatingsMadeByCustomerRequest) (*rating_pb.GetAllRatingsResponse, error) {
	ratings, err := h.ratingService.GetAllRatingsMadeByCustomer(request.Id)

	if err != nil {
		return &rating_pb.GetAllRatingsResponse{
			Ratings: nil,
			RequestResult: &common_pb.RequestResult{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	ratingsResponse := make([]*rating_pb.Rating, 0, len(ratings))

	for _, rating := range ratings {
		ratingsResponse = append(ratingsResponse, &rating_pb.Rating{
			Id:              rating.Id,
			AccommodationId: rating.AccommodationId,
			UserId:          rating.UserId,
			Rating:          rating.Rating,
			HostId:          rating.HostId,
		})
	}

	return &rating_pb.GetAllRatingsResponse{
		Ratings: ratingsResponse,
		RequestResult: &common_pb.RequestResult{
			Code:    200,
			Message: "Successfully retrieved all ratings made by customer.",
		},
	}, nil
}

func (h *RatingHandler) GetAllRatingsForHost(ctx context.Context, request *rating_pb.GetAllRatingsForHostRequest) (*rating_pb.GetAllRatingsResponse, error) {
	ratings, err := h.ratingService.GetAllRatingsForHost(request.Id)

	if err != nil {
		return &rating_pb.GetAllRatingsResponse{
			Ratings: nil,
			RequestResult: &common_pb.RequestResult{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	ratingsResponse := make([]*rating_pb.Rating, 0, len(ratings))

	for _, rating := range ratings {
		ratingsResponse = append(ratingsResponse, &rating_pb.Rating{
			Id:              rating.Id,
			AccommodationId: rating.AccommodationId,
			UserId:          rating.UserId,
			Rating:          rating.Rating,
			HostId:          rating.HostId,
		})
	}

	return &rating_pb.GetAllRatingsResponse{
		Ratings: ratingsResponse,
		RequestResult: &common_pb.RequestResult{
			Code:    200,
			Message: "Successfully retrieved all ratings for host.",
		},
	}, nil
}

func (h *RatingHandler) RateHost(ctx context.Context, request *rating_pb.RateHostRequest) (*common_pb.RequestResult, error) {
	rating := &domain.Rating{
		UserId: request.UserId,
		Rating: request.Rating,
		HostId: request.Id,
	}

	err := h.ratingService.CreateRating(rating)

	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "Successfully rated host.",
	}, nil
}

func (h *RatingHandler) RateAccommodation(ctx context.Context, request *rating_pb.RateAccommodationRequest) (*common_pb.RequestResult, error) {
	rating := &domain.Rating{
		UserId:          request.UserId,
		Rating:          request.Rating,
		AccommodationId: request.Id,
	}

	err := h.ratingService.CreateRating(rating)

	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "Successfully rated accommodation.",
	}, nil
}

func (h *RatingHandler) UpdateRating(ctx context.Context, request *rating_pb.UpdateRatingRequest) (*common_pb.RequestResult, error) {
	err := h.ratingService.UpdateRating(request.Id, request.Rating)

	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "Successfully updated rating.",
	}, nil
}

func (h *RatingHandler) DeleteRating(ctx context.Context, request *rating_pb.DeleteRatingRequest) (*common_pb.RequestResult, error) {
	err := h.ratingService.DeleteRating(request.Id)

	if err != nil {
		return &common_pb.RequestResult{
			Code:    500,
			Message: err.Error(),
		}, nil
	}

	return &common_pb.RequestResult{
		Code:    200,
		Message: "Successfully deleted rating.",
	}, nil
}

func (h *RatingHandler) GetAverageRatingForHost(ctx context.Context, request *rating_pb.GetAverageRatingForHostRequest) (*rating_pb.GetAverageRatingForHostResponse, error) {
	averageRating, err := h.ratingService.GetAverageRatingForHost(request.Id)

	if err != nil {
		return &rating_pb.GetAverageRatingForHostResponse{
			Rating: 0,
			RequestResult: &common_pb.RequestResult{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &rating_pb.GetAverageRatingForHostResponse{
		Rating: averageRating,
		RequestResult: &common_pb.RequestResult{
			Code:    200,
			Message: "Successfully retrieved average rating for host.",
		},
	}, nil
}

func (h *RatingHandler) GetRatingForAccommodation(ctx context.Context, request *rating_pb.GetRatingForAccommodationRequest) (*rating_pb.GetRatingForAccommodationResponse, error) {
	rating, err := h.ratingService.GetRatingForAccommodation(request.Id)

	if err != nil {
		return &rating_pb.GetRatingForAccommodationResponse{
			Rating: 0,
			RequestResult: &common_pb.RequestResult{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	return &rating_pb.GetRatingForAccommodationResponse{
		Rating: rating,
		RequestResult: &common_pb.RequestResult{
			Code:    200,
			Message: "Successfully retrieved rating for accommodation.",
		},
	}, nil
}
