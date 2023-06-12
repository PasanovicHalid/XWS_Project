package presentation

import (
	"context"

	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
	user_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/UserService/domain"
)

type UserHandler struct {
	user_pb.UnimplementedUserServiceServer
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *user_pb.CreateUserRequest) (response *user_pb.CreateUserResponse, err error) {
	user := &domain.User{
		IdentityId:  request.IdentityId,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	}

	err = handler.userService.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return &user_pb.CreateUserResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *user_pb.UpdateUserRequest) (response *user_pb.UpdateUserResponse, err error) {
	user := &domain.User{
		IdentityId:  request.IdentityId,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	}

	err = handler.userService.UpdateUser(user)

	if err != nil {
		if err == persistance.ErrorUserNotFound {
			return &user_pb.UpdateUserResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: err.Error(),
				},
			}, nil
		}

		return nil, err
	}

	return &user_pb.UpdateUserResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *UserHandler) GetUserById(ctx context.Context, request *user_pb.GetUserByIdRequest) (response *user_pb.GetUserByIdResponse, err error) {
	user, err := handler.userService.GetUserById(request.Id)

	if err != nil {
		if err == persistance.ErrorUserNotFound {
			return &user_pb.GetUserByIdResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: err.Error(),
				},
			}, nil
		}

		return nil, err
	}

	return &user_pb.GetUserByIdResponse{
		User: &user_pb.User{
			IdentityId:  user.IdentityId,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
		},
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *UserHandler) GetAllUsers(ctx context.Context, request *user_pb.GetAllUsersRequest) (response *user_pb.GetAllUsersResponse, err error) {
	users, err := handler.userService.GetAllUsersByIdList(request.Ids)

	if err != nil {
		return nil, err
	}

	var usersResponse []*user_pb.User

	for _, user := range users {
		usersResponse = append(usersResponse, &user_pb.User{
			IdentityId:  user.IdentityId,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
		})
	}

	return &user_pb.GetAllUsersResponse{
		Users: usersResponse,
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}
