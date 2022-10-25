package user

import (
	"context"

	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=user -source=./contracts.go

type IService interface {
	CreateUser(ctx context.Context, user *internal.User) (apiErr *apierr.Message)
	GetUserByEmail(ctx context.Context, userEmail string) (user *internal.User, apiErr *apierr.Message)
}
