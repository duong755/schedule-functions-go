package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"schedule.functions/database"
	"schedule.functions/models"
	"schedule.functions/utils"
)

// SchedulesHandler returns some records from Schedule collection
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	studentId := request.URL.Query().Get("studentId")

	regexpStudentId := regexp.MustCompile(`^\d{8}$`)

	if !regexpStudentId.MatchString(studentId) {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(400)
		errorResponse := &utils.ErrorResponse{Message: "Invalid student id"}
		jsonResult, _ := json.Marshal(errorResponse)
		fmt.Println(string(jsonResult))
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}

	dbcontext, client := database.Client()

	database := client.Database("uet")
	scheduleCollection := database.Collection("schedule")

	matchStage := primitive.D{
		{Key: "$match", Value: primitive.D{
			{Key: "studentId", Value: studentId},
		}},
	}
	// lookupStage := primitive.D{
		// {Key: "$lookup", Value: primitive.D{
			// {Key: "", Value: ""},
		// }},
	// }
	scheduleCursor, _ := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{matchStage})

	scheduleRecords := make([]models.Schedule, 0)
	scheduleCursor.All(context.TODO(), &scheduleRecords)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	jsonResult, _ := json.MarshalIndent(scheduleRecords, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
