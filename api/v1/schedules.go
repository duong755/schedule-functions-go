package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"schedule.functions/models"
)

var CONNECTION_STRING string = os.Getenv("CONNECTION_STRING")

// SchedulesHandler returns some records from Schedule collection
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		panic(err)
	}

	database := client.Database("TimetableUET")
	scheduleCollection := database.Collection("Schedule")
	filter := bson.M{"MaSV": "17020191"}
	scheduleCursor, _ := scheduleCollection.Find(ctx, filter)

	scheduleRecords := make([]models.Schedule, 0)
	scheduleCursor.All(ctx, &scheduleRecords)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	jsonResult, _ := json.MarshalIndent(scheduleRecords, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
