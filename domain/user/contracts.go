package user

import (
	"context"

	"secrets-api/domain"
	apierr "secrets-api/infra/errors"
)

type IService interface {
	CreateUser(ctx context.Context, user *domain.User) (apiErr *apierr.Message)
	GetUserByEmail(ctx context.Context, userEmail string) (user *domain.User, apiErr *apierr.Message)
}
