package secret

import (
	"context"

	apierr "github.com/sstalschus/secrets-api/infra/errors"
	"github.com/sstalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=secret -source=./contracts.go

type IService interface {
	CreateSecret(ctx context.Context, secret *internal.Secret, userID string) *apierr.Message
	GetSecrets(ctx context.Context, userID string) *[]internal.Secret
	GetSecret(ctx context.Context, secretID, userID string) (*internal.Secret, *apierr.Message)
}
