package user

import (
	"context"
	"fmt"
	"time"

	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/hash"

	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/sstalschus/secrets-api/internal"
)

type Service struct {
	logger     log.Provider
	repository user_repo.IRepository
	apiErr     apierr.Provider
	auth       hash.Provider
}

func NewService(
	logger log.Provider,
	repository user_repo.IRepository,
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

func (s Service) CreateUser(ctx context.Context, user *internal.User) (apiErr *apierr.Message) {
	userAlreadyExists, _ := s.repository.FindUserByEmail(ctx, user.Email)
	if userAlreadyExists != nil {
		return s.apiErr.BadRequest("User already exists", fmt.Errorf(""))
	}

	user.Id = s.repository.GenerateID()

	apiErr = s.encryptPassword(user)
	if apiErr != nil {
		return apiErr
	}

	s.setTimestamps(user)

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return s.apiErr.BadRequest("Error in register process", err)
	}

	s.logger.Info(ctx, fmt.Sprintf("User created: %s", id), log.Body{})

	return apiErr
}

func (s Service) encryptPassword(user *internal.User) *apierr.Message {
	passwordHash, err := s.auth.Encrypt(user.Password, user.Id.Hex())
	if err != nil {
		return s.apiErr.InternalServerError(err)
	}

	user.Password = string(passwordHash)
	return nil
}

func (s Service) setTimestamps(user *internal.User) {
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()
}

func (s Service) GetUser(ctx context.Context, userID string) (user *internal.User, apiErr *apierr.Message) {
	user, _ = s.repository.FindUserByID(ctx, userID)

	if user == nil {
		apiErr = s.apiErr.BadRequest("User don't exists", fmt.Errorf(""))
	}
	return user, apiErr
}
