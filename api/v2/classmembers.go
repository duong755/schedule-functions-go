package v2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	classes := []modelsV2.Class{}
	classMatchStage := primitive.D{
		{Key: "$match", Value: primitive.D{
			{Key: "classId", Value: classId},
		}},
	}
	classGroupStage := primitive.D{
		{Key: "$group", Value: primitive.D{
			{Key: "_id", Value: "$classId"},
			{Key: "id", Value: primitive.D{
				{Key: "$first", Value: "$_id"},
			}},
			{Key: "classId", Value: primitive.D{
				{Key: "$first", Value: "$classId"},
			}},
			{Key: "subjectId", Value: primitive.D{
				{Key: "$first", Value: "$subjectId"},
			}},
			{Key: "subjectName", Value: primitive.D{
				{Key: "$first", Value: "$subjectName"},
			}},
			{Key: "credit", Value: primitive.D{
				{Key: "$first", Value: "$credit"},
			}},
			{Key: "groups", Value: primitive.D{
				{Key: "$addToSet", Value: primitive.D{
					{Key: "session", Value: "$session"},
					{Key: "weekDay", Value: "$weekDay"},
					{Key: "place", Value: "$place"},
					{Key: "teacher", Value: "$teacher"},
					{Key: "note", Value: "$note"},
				}},
			}},
		}},
	}
	classProjectStage := primitive.D{
		{Key: "$project", Value: primitive.D{
			{Key: "_id", Value: "$id"},
			{Key: "classId", Value: "$classId"},
			{Key: "subjectId", Value: "$subjectId"},
			{Key: "subjectName", Value: "$subjectName"},
			{Key: "credit", Value: "$credit"},
			{Key: "groups", Value: "$groups"},
		}},
	}

	classCursor, classAggregateErr := classCollection.Aggregate(dbcontext, mongo.Pipeline{classMatchStage, classGroupStage, classProjectStage})

	if classAggregateErr != nil {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)
		errorResponse := &utils.ErrorResponse{Message: "Error occured while finding class"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}

	classCursor.All(dbcontext, &classes)

	if len(classes) == 0 {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusNotFound)
		errorResponse := &utils.ErrorResponse{Message: "Class does not exist"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}
	class := classes[0]

	studentMatchStage := primitive.D{
		{Key: "$match", Value: primitive.D{
			{Key: "classId", Value: classId},
		}},
	}
	studentGroupStage := primitive.D{
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
	studentCursor, studentAggregateErr := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{studentMatchStage, studentGroupStage})

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
