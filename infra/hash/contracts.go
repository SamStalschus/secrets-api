package hash

//go:generate mockgen -destination=./mocks.go -package=hash -source=./contracts.go

type Provider interface {
	Encrypt(plaintext, key string) ([]byte, error)
	Decrypt(ciphertext, key string) ([]byte, error)
	Check(ciphertext, plaintext, key string) error
	NewJwt(string) (string, error)
	ValidateJwt(token string) (string, error)
}
