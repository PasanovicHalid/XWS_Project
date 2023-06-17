package authentification

import (
	"context"
	"crypto/x509"
	"time"

	authentification_interface "github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/application/common/interfaces/persistance"
	"github.com/PasanovicHalid/XWS_Project/BookingService/AuthentificationService/domain"
	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	keyRepository persistance.IKeyRepository
}

func NewJwtService(keyRepository persistance.IKeyRepository) *JwtService {
	return &JwtService{
		keyRepository: keyRepository,
	}
}

func (jwtService *JwtService) GenerateToken(id string, username string, role string, apiKey string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	claims := &authentification_interface.SignedDetails{
		Id:       id,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Hour * time.Duration(24)).Unix(),
			Issuer:    "test",
		},
		ApiKey: apiKey,
	}

	keyPair, err := jwtService.keyRepository.GetKeyPair(&ctx)
	if err != nil {
		return "", err
	}

	signingPemBlock := domain.ConvertKeyStringToPemBlock(keyPair.PrivateKey)
	privateKey, _ := x509.ParsePKCS1PrivateKey(signingPemBlock.Bytes)

	return jwt.NewWithClaims(jwt.SigningMethodRS512, claims).SignedString(privateKey)
}

func (jwtService *JwtService) ValidateToken(token string) (claims *authentification_interface.SignedDetails, result int) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	keyPair, err := jwtService.keyRepository.GetKeyPair(&ctx)
	if err != nil {
		result = authentification_interface.Token_Failed
		return
	}

	signingPemBlock := domain.ConvertKeyStringToPemBlock(keyPair.PublicKey)

	tokenDetails, err := jwt.ParseWithClaims(token,
		&authentification_interface.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return x509.ParsePKCS1PublicKey(signingPemBlock.Bytes)
		})

	if err != nil {
		result = authentification_interface.Token_Failed
		return
	}

	claims, ok := tokenDetails.Claims.(*authentification_interface.SignedDetails)
	if !ok {
		result = authentification_interface.Token_Invalid
		return
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		result = authentification_interface.Token_Expired
		return
	}

	return claims, authentification_interface.Token_Valid
}
