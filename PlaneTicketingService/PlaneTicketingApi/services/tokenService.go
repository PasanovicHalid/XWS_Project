package services

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Username  string
	FirstName string
	LastName  string
	Uid       string
	jwt.StandardClaims
}

var SECRET_KEY string

func GenerateToken(username string, firstName string, lastName string, uid string) (token string, err error) {
	claims := &SignedDetails{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
}
