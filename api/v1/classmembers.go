package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"schedule.functions/database"
	modelsV1 "schedule.functions/models/v1"
	"schedule.functions/utils"
)

// ClassMembersHandler returns class info and list of students in it
func ClassMembersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	classId := request.URL.Query().Get("classId")
	classId = strings.Trim(classId, " ")
	classId = strings.ToUpper(classId)

	dbcontext, client := database.Client()
	db := client.Database("uet")
	classCollection := db.Collection("class")
	scheduleCollection := db.Collection("schedule")

	filterClass := bson.D{
		{Key: "classId", Value: classId},
	}

	class := modelsV1.Class{}

	if errFindClass := classCollection.FindOne(dbcontext, filterClass).Decode(&class); errFindClass != nil {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusBadRequest)
		errorResponse := &utils.ErrorResponse{Message: "ClassId does not exist"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}

	matchStage := primitive.D{
		{Key: "$match", Value: primitive.D{
			{Key: "classId", Value: classId},
		}},
	}
	groupStage := primitive.D{
		{Key: "$group", Value: primitive.D{
			{Key: "_id", Value: "$_id"},
			{Key: "studentId", Value: primitive.D{{Key: "$first", Value: "$studentId"}}},
			{Key: "studentName", Value: primitive.D{{Key: "$first", Value: "$studentName"}}},
			{Key: "studentNote", Value: primitive.D{{Key: "$first", Value: "$studentNote"}}},
		}},
	}
	scheduleCursor, _ := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{matchStage, groupStage})

	scheduleRecords := []modelsV1.Schedule{}
	scheduleCursor.All(dbcontext, &scheduleRecords)

	class.Students = scheduleRecords
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	jsonResult, _ := json.MarshalIndent(class, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
