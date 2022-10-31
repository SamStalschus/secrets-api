package user_repo

import (
	"context"

	"github.com/SamStalschus/secrets-api/infra/mongodb"
	"github.com/SamStalschus/secrets-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r Repository) CreateUser(ctx context.Context, user *internal.User) (string, error) {
	res, err := r.repository.InsertOne(ctx, collection, user)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (r Repository) FindUserByEmail(ctx context.Context, email string) (user *internal.User, err error) {
	projection := options.FindOne().SetProjection(bson.M{"password": 0, "secrets": 0})

	err = r.repository.FindOne(ctx, collection, bson.M{"email": email}, projection).Decode(&user)

	return user, err
}

func (r Repository) FindWithPasswordByEmail(ctx context.Context, email string) (user *internal.User, err error) {
	projection := options.FindOne().SetProjection(bson.M{"secrets": 0})
	err = r.repository.FindOne(ctx, collection, bson.M{"email": email}, projection).Decode(&user)

	return user, err
}

func (r Repository) FindUserByID(ctx context.Context, id string) (user *internal.User, err error) {
	projection := options.FindOne().SetProjection(bson.M{"password": 0})

	objectID, _ := primitive.ObjectIDFromHex(id)

	err = r.repository.FindOne(ctx, collection, bson.M{"_id": objectID}, projection).Decode(&user)

	return user, err
}

func (r Repository) GenerateID() primitive.ObjectID {
	return primitive.NewObjectID()
}
