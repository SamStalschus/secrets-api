package user_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/SamStalschus/secrets-api/internal"
)

//go:generate mockgen -destination=./mocks.go -package=user_repo -source=./contracts.go

type IRepository interface {
	CreateUser(ctx context.Context, user *internal.User) (string, error)
	FindUserByEmail(ctx context.Context, email string) (user *internal.User, err error)
	FindWithPasswordByEmail(ctx context.Context, email string) (user *internal.User, err error)
	FindUserByID(ctx context.Context, id string) (user *internal.User, err error)
	GenerateID() primitive.ObjectID
}
