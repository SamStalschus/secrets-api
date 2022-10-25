package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"secrets-api/infra/log"
)

func GetConnection(logger log.Provider, uri string) (*mongo.Client, context.Context) {
	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal(ctx, "Error to connect in database")
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal(ctx, "Error to connect in database")
	}

	return client, ctx
}
