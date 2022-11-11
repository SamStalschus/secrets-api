package secret

import (
	"context"
	"fmt"
	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/hash"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/infra/mongodb/secret_repo"
	"github.com/sstalschus/secrets-api/internal"
	"time"
)

type Service struct {
	logger     log.Provider
	repository secret_repo.IRepository
	apiErr     apierr.Provider
	auth       hash.Provider
}

func NewService(
	logger log.Provider,
	repository secret_repo.IRepository,
	apiErr apierr.Provider,
	auth hash.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
		auth:       auth,
	}
}

func (s Service) CreateSecret(ctx context.Context, secret *internal.Secret, userID string) (apiErr *apierr.Message) {
	secret.Id = s.repository.GenerateID()
	keyHash, err := s.auth.Encrypt(secret.Key, secret.Id.Hex())
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to cipher secret key %s", userID), log.Body{})
		return s.apiErr.BadRequest("Error to create secret", err)
	}

	valueHash, err := s.auth.Encrypt(secret.Value, secret.Id.Hex())
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to cipher secret value %s", userID), log.Body{})
		return s.apiErr.BadRequest("Error to create secret", err)
	}

	secret.Key = string(keyHash)
	secret.Value = string(valueHash)
	s.setTimestamps(secret)

	err = s.repository.CreateSecret(ctx, secret, userID)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to create one secret by user %s", userID), log.Body{})
		return s.apiErr.BadRequest("Error to create secret", err)
	}

	return apiErr
}

func (s Service) setTimestamps(secret *internal.Secret) {
	secret.CreatedAt = time.Now()
	secret.UpdatedAt = time.Now()
}

func (s Service) GetSecrets(ctx context.Context, userID string) *[]internal.Secret {
	secrets := s.repository.FindAllByUserId(ctx, userID)
	return s.decryptKeys(&secrets)
}

func (s Service) GetSecret(ctx context.Context, secretID, userID string) (*internal.Secret, *apierr.Message) {
	secret, err := s.repository.FindSecretByID(ctx, secretID)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to get secret value by user %s", userID), log.Body{})
		return nil, s.apiErr.BadRequest("Error to get secret value", err)
	}

	keyValue, err := s.auth.Decrypt(secret.Key, secret.Id.Hex())
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to decrypt secret key by user %s", userID), log.Body{})
		return nil, s.apiErr.BadRequest("Error to get secret value", err)
	}

	value, err := s.auth.Decrypt(secret.Value, secret.Id.Hex())
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to decrypt secret value by user %s", userID), log.Body{})
		return nil, s.apiErr.BadRequest("Error to get secret value", err)
	}

	secret.Key = string(keyValue)
	secret.Value = string(value)
	secret.UpdatedAt = time.Now()

	return secret, nil
}

func (s Service) decryptKeys(doDecrypt *[]internal.Secret) *[]internal.Secret {
	var secrets []internal.Secret

	for _, secret := range *doDecrypt {
		key, _ := s.auth.Decrypt(secret.Key, secret.Id.Hex())
		secret.Key = string(key)
		secrets = append(secrets, secret)
	}
	return &secrets
}
