package user

import (
	"context"
	"fmt"
	"github.com/sstalschus/secrets-api/infra/cache"
	"time"

	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/hash"

	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/sstalschus/secrets-api/internal"
)

const (
	statusOfUser = "USER_STATUS"
	statusTTL    = 43200
)

type Service struct {
	apiErr     apierr.Provider
	auth       hash.Provider
	cache      cache.Provider
	logger     log.Provider
	repository user_repo.IRepository
}

func NewService(
	logger log.Provider,
	repository user_repo.IRepository,
	apiErr apierr.Provider,
	auth hash.Provider,
	cache cache.Provider,
) Service {
	return Service{
		logger:     logger,
		repository: repository,
		apiErr:     apiErr,
		auth:       auth,
		cache:      cache,
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

	s.setCreatedStatus(user)

	id, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return s.apiErr.BadRequest("Error in register process", err)
	}

	s.logger.Info(ctx, fmt.Sprintf("User created: %s", id), log.Body{})

	return apiErr
}

func (s Service) GetUser(ctx context.Context, userID string) (user *internal.User, apiErr *apierr.Message) {
	user, _ = s.repository.FindUserByID(ctx, userID)

	if user == nil {
		apiErr = s.apiErr.BadRequest("User don't exists", fmt.Errorf(""))
	}
	return user, apiErr
}

func (s Service) BlockUserBySuspect(ctx context.Context, user *internal.User) {
	s.setBlockedStatus(user)

	if err := s.repository.UpdateStatus(ctx, user); err != nil {
		s.logger.Error(ctx, fmt.Sprintf("Error to blocked a user: %s error: %s", user.Id.Hex(), err))
	}

	s.updateStatusInCache(ctx, user)
}

func (s Service) FindWithPasswordByEmail(ctx context.Context, email string) (*internal.User, error) {
	return s.repository.FindWithPasswordByEmail(ctx, email)
}

func (s Service) IsValidUser(ctx context.Context, email string) bool {
	user, _ := s.repository.FindUserByEmail(ctx, email)
	s.updateStatusInCache(ctx, user)

	if user.Status == internal.BlockedStatus || user.Status == internal.CancelledStatus {
		return false
	}

	return true
}

func (s Service) updateStatusInCache(ctx context.Context, user *internal.User) {
	cached := s.cache.GetMap(ctx, user.Id.Hex())

	if cached == nil {
		cached = make(map[string]string)
	}

	cached[statusOfUser] = user.Status

	s.cache.SetMap(ctx, user.Id.Hex(), cached, statusTTL)
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

func (s Service) setCreatedStatus(user *internal.User) {
	user.Status = internal.ActiveStatus
	user.StatusDetail = internal.CreatedStatusDetail
}

func (s Service) setBlockedStatus(user *internal.User) {
	user.Status = internal.BlockedStatus
	user.StatusDetail = internal.BySuspectStatusDetail
}
