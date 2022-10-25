package user_repo

import (
	"context"

	"github.com/SamStalschus/secrets-api/domain"
)

//go:generate mockgen -destination=./mocks.go -package=user_repo -source=./contracts.go

type IRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	FindUserByEmail(ctx context.Context, email string) (user *domain.User, err error)
}
