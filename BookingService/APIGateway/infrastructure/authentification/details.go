package authentification

import "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Id       string
	Username string
	Role     string
	jwt.StandardClaims
}

const (
	Token_Valid = iota
	Token_Invalid
	Token_Expired
	Token_Failed
)
