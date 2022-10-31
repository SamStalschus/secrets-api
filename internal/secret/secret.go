package secret

import (
	"context"
	"fmt"
	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/infra/hash"
	"github.com/SamStalschus/secrets-api/infra/log"
	"github.com/SamStalschus/secrets-api/infra/mongodb/secret_repo"
	"github.com/SamStalschus/secrets-api/internal"
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

	err = s.repository.CreateSecret(ctx, secret, userID)
	if err != nil {
		s.logger.Info(ctx, fmt.Sprintf("Error to create one secret by user %s", userID), log.Body{})
		return s.apiErr.BadRequest("Error to create secret", err)
	}

	return apiErr
}

func (s Service) GetSecrets(ctx context.Context, userID string) *[]internal.Secret {
	secrets := s.repository.FindAllByUserId(ctx, userID)
	return &secrets
}
