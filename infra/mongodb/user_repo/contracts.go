package user_repo

import (
	"context"

	"secrets-api/domain"
)

type IRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (string, error)
	FindUserByEmail(ctx context.Context, email string) (user *domain.User, err error)
}
