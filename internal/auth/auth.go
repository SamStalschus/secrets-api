package auth

import (
	"context"
	"fmt"
	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/hash"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/infra/mongodb/user_repo"
	"github.com/sstalschus/secrets-api/internal"
)

type Service struct {
	apiErr     apierr.Provider
	auth       hash.Provider
	logger     log.Provider
	repository user_repo.IRepository
}

func NewService(
	apiErr apierr.Provider,
	auth hash.Provider,
	logger log.Provider,
	repository user_repo.IRepository,
) Service {
	return Service{
		apiErr:     apiErr,
		auth:       auth,
		logger:     logger,
		repository: repository,
	}
}

func (s Service) GenToken(ctx context.Context, authUser *internal.AuthUser, ip string) (*internal.Token, *apierr.Message) {
	user, err := s.repository.FindWithPasswordByEmail(ctx, authUser.Email)
	if err != nil {
		return nil, s.apiErr.BadRequest("Email or password incorrect.", fmt.Errorf("error in find password"))
	}

	err = s.auth.Check(user.Password, authUser.Password, user.Id.Hex())
	if err != nil {
		return nil, s.apiErr.BadRequest("Email or password incorrect.", fmt.Errorf("error in check password"))
	}

	newJwt, err := s.auth.NewJwt(user.Id.Hex())
	if err != nil {
		return nil, s.apiErr.InternalServerError(fmt.Errorf("error in generate jwt process"))
	}

	s.logger.Info(ctx, fmt.Sprintf("Token generated for user %s and IP %s", user.Id.Hex(), ip), log.Body{})

	return &internal.Token{Token: newJwt, Email: user.Email}, nil
}
