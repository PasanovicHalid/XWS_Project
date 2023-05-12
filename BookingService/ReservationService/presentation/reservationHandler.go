package presentation

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/ReservationService/domain"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	reservation_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/reservation_service"
)

type ReservationHandler struct {
	reservation_pb.UnimplementedReservationServiceServer
	reservationService *application.ReservationService
}

func NewReservationHandler(reservationService *application.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}

func (handler *ReservationHandler) CreateReservation(ctx context.Context, request *reservation_pb.CreateReservationRequest) (response *reservation_pb.CreateReservationResponse, err error) {
	// startTime, err := ptypes.Timestamp(request.StartDateTimeUtc)
	// log.Print("PPPPPPPPPPPPPPPPPPPPPPPPPPPP")
	// log.Print(request.Id)
	// if err != nil {
	// 	//return nil, err
	// }

	// endTime, err := ptypes.Timestamp(request.EndDateTimeUtc)
	// if err != nil {
	// 	//return nil, err
	// }

	reservation := &domain.Reservation{
		Id:                   "1",
		AccommodationOfferId: "request.AccommodationOfferId",
		CustomerId:           "request.CustomerId",
		Status:               0,
		NumberOfGuests:       2,
		StartDateTimeUTC:     time.Now(),
		EndDateTimeUTC:       time.Now(),
	}

	err = handler.reservationService.CreateReservation(reservation)

	if err != nil {
		return nil, err
	}

	return &reservation_pb.CreateReservationResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}
