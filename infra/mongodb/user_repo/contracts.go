package user_repo

import (
	"context"

	"github.com/SamStalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=user_repo -source=./contracts.go

type IRepository interface {
	CreateUser(ctx context.Context, user *internal.User) (string, error)
	FindUserByEmail(ctx context.Context, email string) (user *internal.User, err error)
}
