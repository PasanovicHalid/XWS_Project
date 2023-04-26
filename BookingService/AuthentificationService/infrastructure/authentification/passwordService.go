package authentification

import "golang.org/x/crypto/bcrypt"

type PasswordService struct {
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (passwordService *PasswordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (passwordService *PasswordService) ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	check := true
	if err != nil {
		check = false
	}
	return check
}
