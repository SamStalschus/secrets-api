package auth

import (
	"context"
	"fmt"
	"github.com/sstalschus/secrets-api/infra/cache"
	"github.com/sstalschus/secrets-api/internal/user"
	"strconv"

	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/infra/hash"
	"github.com/sstalschus/secrets-api/infra/log"
	"github.com/sstalschus/secrets-api/internal"
)

const incorrectPassword = "INCORRECT_PASSWORD"

type Service struct {
	apiErr      apierr.Provider
	auth        hash.Provider
	cache       cache.Provider
	logger      log.Provider
	userService user.IService
}

func NewService(
	apiErr apierr.Provider,
	auth hash.Provider,
	cache cache.Provider,
	logger log.Provider,
	userService user.IService,
) Service {
	return Service{
		apiErr:      apiErr,
		auth:        auth,
		cache:       cache,
		logger:      logger,
		userService: userService,
	}
}

func (s Service) GenToken(ctx context.Context, authUser *internal.AuthUser, ip string) (*internal.Token, *apierr.Message) {
	if valid := s.userService.IsValidUser(ctx, authUser.Email); !valid {
		return nil, s.apiErr.Blocked()
	}

	user, err := s.userService.FindWithPasswordByEmail(ctx, authUser.Email)
	if err != nil {
		return nil, s.apiErr.BadRequest("Email or password incorrect.", fmt.Errorf("error in find password"))
	}

	err = s.auth.Check(user.Password, authUser.Password, user.Id.Hex())
	if err != nil {
		if suspect := s.isFraudSuspect(ctx, user.Email); suspect {
			s.userService.BlockUserBySuspect(ctx, user)
		}

		return nil, s.apiErr.BadRequest("Email or password incorrect.", fmt.Errorf("error in check password"))
	}

	newJwt, err := s.auth.NewJwt(user.Id.Hex())
	if err != nil {
		return nil, s.apiErr.InternalServerError(fmt.Errorf("error in generate jwt process"))
	}

	s.logger.Info(ctx, fmt.Sprintf("Token generated for user %s and IP %s", user.Id.Hex(), ip), log.Body{})

	return &internal.Token{Token: newJwt, Email: user.Email}, nil
}

func (s Service) isFraudSuspect(ctx context.Context, userEmail string) bool {
	userStatus := s.cache.GetMap(ctx, userEmail)

	finalStatus := s.incrementKey(userStatus, incorrectPassword)

	if number, _ := strconv.Atoi(finalStatus[incorrectPassword]); number > 5 {
		s.logger.Error(ctx, fmt.Sprintf("User %s suspected of fraud", userEmail))

		return true
	}

	s.cache.SetMap(ctx, userEmail, finalStatus, 0)

	return false
}

func (s Service) incrementKey(status map[string]string, key string) map[string]string {
	if status != nil {
		number, _ := strconv.Atoi(status[key])
		number++
		status[key] = strconv.Itoa(number)
		return status
	}

	final := make(map[string]string)
	final[key] = "1"

	return final
}
