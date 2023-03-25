package services

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Username  string
	FirstName string
	LastName  string
	Role      string
	Uid       string
	jwt.StandardClaims
}

const (
	Token_Valid   = iota
	Token_Invalid = iota
	Token_Expired = iota
	Token_Failed  = iota
)

var SECRET_KEY string

func GenerateToken(username string, firstName string, lastName string, role string, uid string) (token string, err error) {
	claims := &SignedDetails{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		Uid:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
}

func ValidateToken(signedToken string) (claims *SignedDetails, result int) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		result = Token_Failed
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		result = Token_Invalid
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		result = Token_Expired
		return
	}

	return claims, Token_Valid
}
