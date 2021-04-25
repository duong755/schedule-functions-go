package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Client() (context.Context, *mongo.Client) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		panic(err)
	}
	return ctx, client
}
