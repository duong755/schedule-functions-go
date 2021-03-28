package v1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SchedulesHandler returns the function name
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	CONNECTION_STRING := os.Getenv("CONNECTION_STRING")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(responseWriter, "Connected to MongoDB")
}
