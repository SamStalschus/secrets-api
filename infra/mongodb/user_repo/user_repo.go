package user_repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"secrets-api/domain"
	"secrets-api/infra/mongodb"
)

type Repository struct {
	repository *mongodb.Repository
}

func NewRepository(
	repository *mongodb.Repository,
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

func (r Repository) FindUserByEmail(ctx context.Context, email string) (error, *mongo.SingleResult) {
	projection := options.FindOne().SetProjection(bson.E{Key: "password", Value: 0})

	res := r.repository.FindOne(ctx, collection, bson.M{"email": email}, projection)

	var result bson.D

	err := res.Decode(&result)

	if err != nil {
		return err, nil
	}
}
