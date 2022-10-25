package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	InsertOne(ctx context.Context, collection string, data any) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, collection string, data any, opts *options.FindOneOptions) *mongo.SingleResult
}
