package authentification

import (
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (service *JwtService) ValidateToken(token string, decryptionKey string) (claims *SignedDetails, result int) {
	signingPemBlock, _ := pem.Decode([]byte(decryptionKey))

	tokenDetails, err := jwt.ParseWithClaims(token,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return x509.ParsePKCS1PublicKey(signingPemBlock.Bytes)
		})

	if err != nil {
		result = Token_Failed
		return
	}

	claims, ok := tokenDetails.Claims.(*SignedDetails)
	if !ok {
		result = Token_Invalid
		return
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		result = Token_Expired
		return
	}

	return claims, Token_Valid
}
