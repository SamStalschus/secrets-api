package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository(
	client *mongo.Client,
) Repository {
	return Repository{
		client: client,
	}
}

const database = "secrets"

func (r Repository) InsertOne(ctx context.Context, collection string, data any) (*mongo.InsertOneResult, error) {
	return r.client.Database(database).Collection(collection).InsertOne(ctx, data)
}

func (r Repository) FindOne(ctx context.Context, collection string, data any, opts *options.FindOneOptions) *mongo.SingleResult {
	return r.client.Database(database).Collection(collection).FindOne(ctx, data, opts)
}
