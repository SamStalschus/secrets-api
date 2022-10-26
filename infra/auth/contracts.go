package auth

//go:generate mockgen -destination=./mocks.go -package=auth -source=./contracts.go

type Provider interface {
	EncryptPassword(password string) ([]byte, error)
	CheckPassword(password, providedPassword string) error
	NewJwt(string) (string, error)
	ValidateJwt(token string) (string, error)
}
