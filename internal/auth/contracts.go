package auth

import (
	"context"

	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=auth -source=./contracts.go

type IService interface {
	GenToken(ctx context.Context, user *internal.AuthUser, ip string) (*internal.Token, *apierr.Message)
}
