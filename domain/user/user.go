package user

import (
	"context"
	"fmt"
	"secrets-api/domain"
	"secrets-api/infra/errors"
	"secrets-api/infra/log"
	"secrets-api/infra/mongodb/user_repo"
)

type Service struct {
	logger     log.Provider
	repository *user_repo.Repository
	apiErr     apierr.Provider
}

func NewService(
	logger log.Provider,
	repository *user_repo.Repository,
	apiErr apierr.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
	}
}

const collection = "users"

func (s Service) CreateUser(ctx context.Context, user *domain.User) (apiErr *apierr.Message) {
	s.repository.FindOne(ctx, collection)

	res, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		apiErr = s.apiErr.BadRequest("Error in register process", err)
	}

	s.logger.Info(ctx, fmt.Sprintf("User created: %s", res), log.Body{})

	return apiErr
}

func (s Service) GetUser(ctx context.Context, userID int) string {
	s.logger.Info(ctx, "user", log.Body{
		"user_id": userID,
	})
	return "fake user" + string(userID)
}
