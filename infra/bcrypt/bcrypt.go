package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c Client) EncryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
