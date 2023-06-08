package presentation

import (
	"context"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/infrastructure/message_queues"
	auth_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	common_pb "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/common"
)

type IdentityHandler struct {
	identityService        *application.IdentityService
	deleteUserOrchestrator *message_queues.DeleteUserOrchestrator
	auth_pb.UnimplementedAuthenticateServiceServer
}

func NewIdentityHandler(identityService *application.IdentityService, deleteUserOrchestrator *message_queues.DeleteUserOrchestrator) *IdentityHandler {
	return &IdentityHandler{
		identityService:        identityService,
		deleteUserOrchestrator: deleteUserOrchestrator,
	}
}

func (handler *IdentityHandler) Register(ctx context.Context, request *auth_pb.RegisterRequest) (*auth_pb.RegisterResponse, error) {
	identity := &domain.Identity{
		Username: request.Username,
		Password: request.Password,
	}

	if request.IsHost {
		identity.Role = domain.Role_Host
	} else {
		identity.Role = domain.Role_Guest
	}

	jwtToken, err := handler.identityService.RegisterIdentity(identity)

	if err != nil {

		if err == persistance.ErrorUsernameInUse {
			return &auth_pb.RegisterResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: err.Error(),
				},
			}, nil
		}

		return nil, err
	}

	return &auth_pb.RegisterResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
		Token: jwtToken,
	}, nil
}

func (handler *IdentityHandler) Login(ctx context.Context, request *auth_pb.LoginRequest) (*auth_pb.LoginResponse, error) {
	jwtToken, err := handler.identityService.Login(request.Username, request.Password)

	if err != nil {
		if err == persistance.ErrorIdentityNotFound || err == persistance.ErrorInvalidPassword {
			return &auth_pb.LoginResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: "Invalid username or password",
				},
			}, nil
		}

		return nil, err
	}

	return &auth_pb.LoginResponse{
		Token: jwtToken,
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *IdentityHandler) GetPublicKey(ctx context.Context, request *auth_pb.GetPublicKeyRequest) (*auth_pb.GetPublicKeyResponse, error) {
	keyPair, err := handler.identityService.GetPublicKey()

	if err != nil {
		return nil, err
	}

	return &auth_pb.GetPublicKeyResponse{
		PublicKey: keyPair.PublicKey,
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *IdentityHandler) ChangePassword(ctx context.Context, request *auth_pb.ChangePasswordRequest) (*auth_pb.ChangePasswordResponse, error) {
	err := handler.identityService.ChangePassword(request.Username, request.OldPassword, request.NewPassword)

	if err != nil {
		if err == persistance.ErrorIdentityNotFound || err == persistance.ErrorInvalidPassword {
			return &auth_pb.ChangePasswordResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: "Invalid username or password",
				},
			}, nil
		}

		return nil, err
	}

	return &auth_pb.ChangePasswordResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *IdentityHandler) ChangeUsername(ctx context.Context, request *auth_pb.ChangeUsernameRequest) (*auth_pb.ChangeUsernameResponse, error) {
	err := handler.identityService.ChangeUsername(request.Username, request.Password, request.NewUsername)

	if err != nil {
		if err == persistance.ErrorIdentityNotFound || err == persistance.ErrorUsernameInUse {
			return &auth_pb.ChangeUsernameResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: "Invalid username",
				},
			}, nil
		}

		if err == persistance.ErrorInvalidPassword {
			return &auth_pb.ChangeUsernameResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: "Invalid password",
				},
			}, nil
		}

		return nil, err
	}

	return &auth_pb.ChangeUsernameResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *IdentityHandler) GetIdentityByUsername(ctx context.Context, request *auth_pb.GetIdentityByUsernameRequest) (*auth_pb.GetIdentityByUsernameResponse, error) {
	identity, err := handler.identityService.FindIdentityByUsername(request.Username)

	if err != nil {
		if err == persistance.ErrorIdentityNotFound {
			return &auth_pb.GetIdentityByUsernameResponse{
				RequestResult: &common_pb.RequestResult{
					Code:    400,
					Message: "Invalid username",
				},
			}, nil
		}
		return nil, err
	}

	return &auth_pb.GetIdentityByUsernameResponse{
		IdentityId: identity.Id.Hex(),
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}

func (handler *IdentityHandler) Remove(ctx context.Context, request *auth_pb.RemoveRequest) (*auth_pb.RemoveResponse, error) {
	// err := handler.identityService.DeleteIdentity(request.IdentityId)

	// if err != nil {
	// 	if err == persistance.ErrorIdentityNotFound {
	// 		return &auth_pb.RemoveResponse{
	// 			RequestResult: &common_pb.RequestResult{
	// 				Code:    400,
	// 				Message: "Invalid identity id",
	// 			},
	// 		}, nil
	// 	}
	// 	return nil, err
	// }

	err := handler.deleteUserOrchestrator.Start()

	if err != nil {
		return &auth_pb.RemoveResponse{
			RequestResult: &common_pb.RequestResult{
				Code:    400,
				Message: "Something failed with orchestrator",
			},
		}, nil
	}

	return &auth_pb.RemoveResponse{
		RequestResult: &common_pb.RequestResult{
			Code: 200,
		},
	}, nil
}
