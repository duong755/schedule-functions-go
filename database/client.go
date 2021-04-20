package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Client() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		panic(err)
	}
	return client
}
