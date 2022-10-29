package secret

import (
	"context"
	apierr "github.com/SamStalschus/secrets-api/infra/errors"
	"github.com/SamStalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=secret -source=./contracts.go

type IService interface {
	CreateSecret(ctx context.Context, secret *internal.Secret, userID string) (apiErr *apierr.Message)
}
