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
	keyRepository      persistance.IKeyRepository
	passwordService    authentification.IPasswordService
	jwtService         authentification.IJwtService
}

func NewIdentityService(identityRepository persistance.IIdentityRepository, keyRepository persistance.IKeyRepository, passwordService authentification.IPasswordService, jwtService authentification.IJwtService) *IdentityService {
	return &IdentityService{
		identityRepository: identityRepository,
		keyRepository:      keyRepository,
		passwordService:    passwordService,
		jwtService:         jwtService,
	}
}

func (service *IdentityService) FindIdentityByEmail(email string) (*domain.Identity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.FindIdentityByEmail(&ctx, email)
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

func (service *IdentityService) DeleteIdentity(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	return service.identityRepository.DeleteIdentity(&ctx, id)
}

func (service *IdentityService) RegisterIdentity(identity *domain.Identity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	exists, err := service.identityRepository.CheckIfEmailExists(&ctx, identity.Email)

	if err != nil {
		return err
	}

	if exists {
		return persistance.ErrorEmailInUse
	}

	identity.Password, err = service.passwordService.HashPassword(identity.Password)

	if err != nil {
		return err
	}

	return service.identityRepository.InsertIdentity(&ctx, identity)
}

func (service *IdentityService) Login(email string, password string) (jwtToken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	identity, err := service.identityRepository.FindIdentityByEmail(&ctx, email)

	if err != nil {
		return jwtToken, err
	}

	if identity == nil {
		return jwtToken, persistance.ErrorIdentityWithEmailDoesntExist
	}

	if !service.passwordService.ComparePasswords(identity.Password, password) {
		return jwtToken, persistance.ErrorInvalidPassword
	}

	jwtToken, err = service.jwtService.GenerateToken(identity.Id.Hex(), identity.Email, identity.Role)

	return jwtToken, nil
}
