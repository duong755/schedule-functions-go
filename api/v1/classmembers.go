package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"schedule.functions/database"
	"schedule.functions/models"
	"schedule.functions/utils"
)

// ClassMembersHandler returns the function name
func ClassMembersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	classId := request.URL.Query().Get("classId")

	dbcontext, client := database.Client()
	db := client.Database("uet")
	classCollection := db.Collection("class")
	scheduleCollection := db.Collection("schedule")

	filterClass := bson.D{
		{Key: "classId", Value: classId},
	}

	class := models.Class{}

	if errFindClass := classCollection.FindOne(dbcontext, filterClass).Decode(&class); errFindClass != nil {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(404)
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
			{Key: "studentId", Value: primitive.D{{ Key: "$first", Value: "$studentId" }}},
			{Key: "studentName", Value: primitive.D{{ Key: "$first", Value: "$studentName" }}},
			{Key: "studentNote", Value: primitive.D{{ Key: "$first", Value: "$studentNote" }}},
		}},
	}
	scheduleCursor, _ := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{matchStage, groupStage})

	scheduleRecords := make([]models.Schedule, 0)
	scheduleCursor.All(context.TODO(), &scheduleRecords)

	class.Students = scheduleRecords
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	jsonResult, _ := json.MarshalIndent(class, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
