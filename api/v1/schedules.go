package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"schedule.functions/database"
	modelsV1 "schedule.functions/models/v1"
	"schedule.functions/utils"
)

// SchedulesHandler returns some records from Schedule collection
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	studentId := request.URL.Query().Get("studentId")
	studentId = strings.Trim(studentId, " ")
	studentId = strings.ToUpper(studentId)
	regexpStudentId := regexp.MustCompile(`^\d{8}$`)

	if !regexpStudentId.MatchString(studentId) {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusBadRequest)
		errorResponse := &utils.ErrorResponse{Message: "Invalid student id"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}

	dbcontext, client := database.Client()

	db := client.Database("uet")
	scheduleCollection := db.Collection("schedule")

	matchStage := primitive.D{
		{Key: "$match", Value: primitive.D{
			{Key: "studentId", Value: studentId},
		}},
	}
	lookupStage := primitive.D{
		{Key: "$lookup", Value: primitive.D{
			{Key: "from", Value: "class"},
			{Key: "localField", Value: "classId"},
			{Key: "foreignField", Value: "classId"},
			{Key: "as", Value: "classes"},
		}},
	}
	scheduleCursor, _ := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{matchStage, lookupStage})

	scheduleRecords := []modelsV1.Schedule{}
	scheduleCursor.All(dbcontext, &scheduleRecords)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	jsonResult, _ := json.MarshalIndent(scheduleRecords, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
