package presentation

import (
	"context"

	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	rec_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/recommendation_service"
)

type RecommendationHandler struct {
	rec_pb.UnimplementedRecommendationServiceServer
}

func NewRecommendationHandler() *RecommendationHandler {
	return &RecommendationHandler{}
}

func (handler *RecommendationHandler) GetRecommendation(ctx context.Context, request *rec_pb.RecommendationRequest) (*common_pb.RequestResult, error) {
	return &common_pb.RequestResult{
		Code:    200,
		Message: "Recommendation sent successfully",
	}, nil
}
