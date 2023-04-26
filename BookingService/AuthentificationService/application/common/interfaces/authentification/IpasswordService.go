package authentification

type IPasswordService interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword string, password string) bool
}
