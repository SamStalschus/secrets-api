package user_repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"secrets-api/domain"
	"secrets-api/infra/mongodb"
)

type Repository struct {
	repository mongodb.IRepository
}

func NewRepository(
	repository mongodb.IRepository,
) Repository {
	return Repository{
		repository: repository,
	}
}

const collection = "users"

func (r Repository) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	res, err := r.repository.InsertOne(ctx, collection, user)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r Repository) FindUserByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	projection := options.FindOne().SetProjection(bson.M{"password": 0})

	err = r.repository.FindOne(ctx, collection, bson.M{"email": email}, projection).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, err
}
