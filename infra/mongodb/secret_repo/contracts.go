package secret_repo

import (
	"context"
	"github.com/SamStalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=secret_repo -source=./contracts.go

type IRepository interface {
	CreateSecret(ctx context.Context, secret *internal.Secret, userID string) error
}
