package application

import (
	"context"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
)

type IdentityService struct {
	identityRepository persistance.IIdentityRepository
	keyService         authentification.IKeyService
	passwordService    authentification.IPasswordService
	jwtService         authentification.IJwtService
}

func NewIdentityService(identityRepository persistance.IIdentityRepository, keyService authentification.IKeyService, passwordService authentification.IPasswordService, jwtService authentification.IJwtService) *IdentityService {
	return &IdentityService{
		identityRepository: identityRepository,
		keyService:         keyService,
		passwordService:    passwordService,
		jwtService:         jwtService,
	}
}

func (service *IdentityService) FindIdentityByUsername(username string) (*domain.Identity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.FindIdentityByUsername(&ctx, username)
}

func (service *IdentityService) FindIdentityById(id string) (*domain.Identity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.FindIdentityById(&ctx, id)
}

func (service *IdentityService) UpdateIdentity(identity *domain.Identity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.UpdateIdentity(&ctx, identity)
}

func (service *IdentityService) DeleteIdentity(id string, sagaTimestamp int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.DeleteIdentity(&ctx, id, sagaTimestamp)
}

func (service *IdentityService) RollbackDeleteIdentity(id string, sagaTimestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.RollbackDeleteIdentity(&ctx, id, sagaTimestamp)
}

func (service *IdentityService) RegisterIdentity(identity *domain.Identity) (jwtToken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	exists, err := service.identityRepository.CheckIfUsernameExists(&ctx, identity.Username)

	if err != nil {
		return
	}

	if exists {
		err = persistance.ErrorUsernameInUse
		return
	}

	identity.Password, err = service.passwordService.HashPassword(identity.Password)

	if err != nil {
		return
	}

	err = service.identityRepository.InsertIdentity(&ctx, identity)

	if err != nil {
		return
	}

	return service.jwtService.GenerateToken(identity.Id.Hex(), identity.Username, identity.Role)
}

func (service *IdentityService) Login(username string, password string) (jwtToken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	identity, err := service.identityRepository.FindIdentityByUsername(&ctx, username)

	if err != nil {
		return jwtToken, err
	}

	if !service.passwordService.ComparePasswords(identity.Password, password) {
		return jwtToken, persistance.ErrorInvalidPassword
	}

	jwtToken, err = service.jwtService.GenerateToken(identity.Id.Hex(), identity.Username, identity.Role)

	if err != nil {
		return jwtToken, err
	}

	return jwtToken, nil
}

func (service *IdentityService) GetPublicKey() (*domain.KeyPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.keyService.RetrieveKeyPair(&ctx)
}

func (service *IdentityService) ChangePassword(username string, oldPassword string, newPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	identity, err := service.identityRepository.FindIdentityByUsername(&ctx, username)

	if err != nil {
		return err
	}

	if !service.passwordService.ComparePasswords(identity.Password, oldPassword) {
		return persistance.ErrorInvalidPassword
	}

	identity.Password, err = service.passwordService.HashPassword(newPassword)

	if err != nil {
		return err
	}

	return service.identityRepository.UpdateIdentity(&ctx, identity)
}

func (service *IdentityService) ChangeUsername(oldUsername string, password string, newUsername string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	identity, err := service.identityRepository.FindIdentityByUsername(&ctx, oldUsername)

	if err != nil {
		return err
	}

	if !service.passwordService.ComparePasswords(identity.Password, password) {
		return persistance.ErrorInvalidPassword
	}

	exists, err := service.identityRepository.CheckIfUsernameExists(&ctx, newUsername)

	if err != nil {
		return err
	}

	if exists {
		return persistance.ErrorUsernameInUse
	}

	identity.Username = newUsername

	return service.identityRepository.UpdateIdentity(&ctx, identity)
}
