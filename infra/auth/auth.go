package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Client struct {
	SecretKey string
}

func NewClient(secretKey string) Client {
	return Client{
		SecretKey: secretKey,
	}
}

func (c Client) EncryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (c Client) CheckPassword(password, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
}

func (c Client) NewJwt(sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "secrets",
		"sub": sub,
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(c.SecretKey))
}

func (c Client) ValidateJwt(token string) (string, error) {
	tkn, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if tkn.Valid && err == nil {
		claims := tkn.Claims.(jwt.MapClaims)
		sub := claims["sub"].(string)
		return sub, err
	}

	return "", err
}
