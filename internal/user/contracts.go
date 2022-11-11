package user

import (
	"context"

	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=user -source=./contracts.go

type IService interface {
	CreateUser(ctx context.Context, user *internal.User) (apiErr *apierr.Message)
	GetUser(ctx context.Context, userID string) (user *internal.User, apiErr *apierr.Message)
}
