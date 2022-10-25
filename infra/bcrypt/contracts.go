package bcrypt

//go:generate mockgen -destination=./mocks.go -package=bcrypt -source=./contracts.go

type Provider interface {
	EncryptPassword(password string) ([]byte, error)
}
