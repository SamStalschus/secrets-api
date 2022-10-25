package user

import (
	"context"
	"fmt"
	"secrets-api/infra/bcrypt"

	"secrets-api/domain"
	"secrets-api/infra/errors"
	"secrets-api/infra/log"
	"secrets-api/infra/mongodb/user_repo"
)

type Service struct {
	logger     log.Provider
	repository user_repo.IRepository
	apiErr     apierr.Provider
	bcrypt     bcrypt.Provider
}

func NewService(
	logger log.Provider,
	repository user_repo.IRepository,
	apiErr apierr.Provider,
	bcrypt bcrypt.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
		bcrypt:     bcrypt,
	}
}

func (s Service) CreateUser(ctx context.Context, user *domain.User) (apiErr *apierr.Message) {
	userAlreadyExists, _ := s.repository.FindUserByEmail(ctx, user.Email)

	if userAlreadyExists != nil {
		return s.apiErr.BadRequest("User already exists", fmt.Errorf(""))
	}

	passwordHash, err := s.bcrypt.EncryptPassword(user.Password)
	if err != nil {
		return s.apiErr.InternalServerError(err)
	}

	user.Password = string(passwordHash)

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return s.apiErr.BadRequest("Error in register process", err)
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
