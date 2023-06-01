package presentation

import (
	"context"

	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	rating_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/rating_service"
)

type RatingHandler struct {
	rating_pb.UnimplementedRatingServiceServer
}

func NewRatingHandler() *RatingHandler {
	return &RatingHandler{}
}

func (handler *RatingHandler) GetRating(context context.Context, request *rating_pb.GetRatingRequest) (*common_pb.RequestResult, error) {
	return &common_pb.RequestResult{
		Code:    200,
		Message: "Rating sent successfully",
	}, nil
}
