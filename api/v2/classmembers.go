package v2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"schedule.functions/database"
	modelsV2 "schedule.functions/models/v2"
	"schedule.functions/utils"
)

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

	class := modelsV2.Class{}

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
			{Key: "studentBirthday", Value: primitive.D{{Key: "$first", Value: "$studentBirthday"}}},
			{Key: "studentCourse", Value: primitive.D{{Key: "$first", Value: "$studentCourse"}}},
			{Key: "classNote", Value: primitive.D{{Key: "$first", Value: "$classNote"}}},
		}},
	}
	studentCursor, studentAggregateErr := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{matchStage, groupStage})

	if studentAggregateErr != nil {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)
		errorResponse := &utils.ErrorResponse{Message: "Error occured while getting class members"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}

	studentRecords := []modelsV2.Student{}
	studentCursor.All(dbcontext, &studentRecords)

	class.Students = studentRecords
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	jsonResult, _ := json.MarshalIndent(class, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
