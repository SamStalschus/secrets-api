package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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

func (c Client) Encrypt(plaintext, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(c.SecretKey))
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := c.getNonce(key)
	return aesGCM.Seal(nil, nonce, []byte(plaintext), nil), nil
}

func (c Client) Decrypt(ciphertext, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(c.SecretKey))
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := c.getNonce(key)
	return aesGCM.Open(nil, nonce, []byte(ciphertext), nil)
}

func (c Client) Check(ciphertext, plaintext, key string) error {
	decrypt, err := c.Decrypt(ciphertext, key)
	if err != nil {
		return err
	}
	if string(decrypt) != plaintext {
		return fmt.Errorf("invalid key")
	}
	return nil
}

func (c Client) NewJwt(sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "secrets",
		"sub": sub,
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 3).Unix(),
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

func (c Client) getNonce(key string) []byte {
	var nonce string
	for i := 0; i < 12; i++ {
		nonce += string(key[i])
	}
	return []byte(nonce)
}
