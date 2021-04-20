package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"schedule.functions/database"
	"schedule.functions/models"
)

// SchedulesHandler returns some records from Schedule collection
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	client := database.Client()

	database := client.Database("TimetableUET")
	scheduleCollection := database.Collection("Schedule")
	filter := bson.M{"MaSV": "17020191"}
	scheduleCursor, _ := scheduleCollection.Find(context.TODO(), filter)

	scheduleRecords := make([]models.Schedule, 0)
	scheduleCursor.All(context.TODO(), &scheduleRecords)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	jsonResult, _ := json.MarshalIndent(scheduleRecords, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
