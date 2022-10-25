package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"secrets-api/domain"
	"secrets-api/infra/errors"
	"secrets-api/infra/log"
	"secrets-api/infra/mongodb/user_repo"
)

type Service struct {
	logger     log.Provider
	repository user_repo.IRepository
	apiErr     apierr.Provider
}

func NewService(
	logger log.Provider,
	repository user_repo.IRepository,
	apiErr apierr.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
	}
}

func (s Service) CreateUser(ctx context.Context, user *domain.User) (apiErr *apierr.Message) {
	userAlreadyExists, _ := s.repository.FindUserByEmail(ctx, user.Email)

	if userAlreadyExists != nil {
		apiErr = s.apiErr.BadRequest("User already exists", fmt.Errorf(""))
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return s.apiErr.InternalServerError(err)
	}

	user.Password = string(passwordHash)

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		apiErr = s.apiErr.BadRequest("Error in register process", err)
	}

	s.logger.Info(ctx, fmt.Sprintf("User created: %s", id), log.Body{})

	return apiErr
}

func (s Service) GetUserByEmail(ctx context.Context, userEmail string) (user *domain.User, apiErr *apierr.Message) {
	user, _ = s.repository.FindUserByEmail(ctx, userEmail)

	if user == nil {
		apiErr = s.apiErr.BadRequest("User don't exists", fmt.Errorf(""))
	}
	return user, apiErr
}
