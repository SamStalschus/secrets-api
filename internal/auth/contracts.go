package auth

import (
	"context"
	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=auth -source=./contracts.go

type IService interface {
	GenToken(ctx context.Context, user *internal.AuthUser, ip string) (*internal.Token, *apierr.Message)
}
