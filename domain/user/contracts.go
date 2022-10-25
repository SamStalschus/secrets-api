package user

import (
	"context"

	"github.com/SamStalschus/secrets-api/domain"
	apierr "github.com/SamStalschus/secrets-api/infra/errors"
)

//go:generate mockgen -destination=./mocks.go -package=user -source=./contracts.go

type IService interface {
	CreateUser(ctx context.Context, user *domain.User) (apiErr *apierr.Message)
	GetUserByEmail(ctx context.Context, userEmail string) (user *domain.User, apiErr *apierr.Message)
}
