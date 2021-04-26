package v2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"schedule.functions/database"
	modelsV2 "schedule.functions/models/v2"
	"schedule.functions/utils"
)

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

	matchStudentIdStage := primitive.D{
		{Key: "$match", Value: primitive.D{
			{Key: "studentId", Value: studentId},
		}},
	}
	groupClassIdAndStudentIdStage := primitive.D{
		{Key: "$group", Value: primitive.D{
			{Key: "_id", Value: primitive.D{
				{Key: "studentId", Value: "$studentId"},
				{Key: "classId", Value: "$classId"},
			}},
			{Key: "id", Value: primitive.D{
				{Key: "$first", Value: "$_id"},
			}},
			{Key: "studentId", Value: primitive.D{
				{Key: "$first", Value: "$studentId"},
			}},
			{Key: "studentName", Value: primitive.D{
				{Key: "$first", Value: "$studentName"},
			}},
			{Key: "studentNote", Value: primitive.D{
				{Key: "$first", Value: "$studentNote"},
			}},
			{Key: "studentCourse", Value: primitive.D{
				{Key: "$first", Value: "$studentCourse"},
			}},
			{Key: "studentBirthday", Value: primitive.D{
				{Key: "$first", Value: "$studentBirthday"},
			}},
			{Key: "classId", Value: primitive.D{
				{Key: "$first", Value: "$classId"},
			}},
			{Key: "classNote", Value: primitive.D{
				{Key: "$first", Value: "$classNote"},
			}},
		}},
	}
	lookupClassStage := primitive.D{
		{Key: "$lookup", Value: primitive.D{
			{Key: "from", Value: "class"},
			{Key: "localField", Value: "classId"},
			{Key: "foreignField", Value: "classId"},
			{Key: "as", Value: "classes"},
		}},
	}
	project1Stage := primitive.D{
		{Key: "$project", Value: primitive.D{
			{Key: "_id", Value: 0},
			{Key: "id", Value: 1},
			{Key: "studentId", Value: 1},
			{Key: "studentName", Value: 1},
			{Key: "studentCourse", Value: 1},
			{Key: "studentBirthday", Value: 1},
			{Key: "studentNote", Value: 1},
			{Key: "classNote", Value: 1},
			{Key: "classes", Value: primitive.D{
				{Key: "$filter", Value: primitive.D{
					{Key: "input", Value: "$classes"},
					{Key: "as", Value: "class"},
					{Key: "cond", Value: primitive.D{
						{Key: "$in", Value: primitive.A{
							"$$class.note",
							primitive.A{"CL", "$classNote"},
						}},
					}},
				}},
			}},
		}},
	}
	project2Stage := primitive.D{
		{Key: "$project", Value: primitive.D{
			{Key: "classNote", Value: 0},
		}},
	}
	unwindStage := primitive.D{
		{Key: "$unwind", Value: primitive.D{
			{Key: "path", Value: "$classes"},
			{Key: "includeArrayIndex", Value: "class"},
			{Key: "preserveNullAndEmptyArrays", Value: false},
		}},
	}
	addFieldsStage := primitive.D{
		{Key: "$addFields", Value: primitive.D{
			{Key: "classes.studentNote", Value: "$studentNote"},
		}},
	}
	group2Stage := primitive.D{
		{Key: "$group", Value: primitive.D{
			{Key: "_id", Value: "$studentId"},
			{Key: "studentId", Value: primitive.D{
				{Key: "$first", Value: "$studentId"},
			}},
			{Key: "studentName", Value: primitive.D{
				{Key: "$first", Value: "$studentName"},
			}},
			{Key: "studentBirthday", Value: primitive.D{
				{Key: "$first", Value: "$studentBirthday"},
			}},
			{Key: "studentCourse", Value: primitive.D{
				{Key: "$first", Value: "$studentCourse"},
			}},
			{Key: "classes", Value: primitive.D{
				{Key: "$addToSet", Value: "$classes"},
			}},
		}},
	}

	scheduleCursor, aggregateErr := scheduleCollection.Aggregate(dbcontext, mongo.Pipeline{
		matchStudentIdStage,
		groupClassIdAndStudentIdStage,
		lookupClassStage,
		project1Stage,
		project2Stage,
		unwindStage,
		addFieldsStage,
		group2Stage,
	})

	if aggregateErr != nil {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)
		errorResponse := &utils.ErrorResponse{Message: "Error occurred while getting schedules"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
	}

	scheduleRecords := []modelsV2.Schedule{}
	scheduleCursor.All(dbcontext, &scheduleRecords)

	if len(scheduleRecords) == 0 {
		responseWriter.Header().Add("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusNotFound)
		errorResponse := &utils.ErrorResponse{Message: "StudentId does not exist"}
		jsonResult, _ := json.MarshalIndent(errorResponse, "", "  ")
		fmt.Fprint(responseWriter, string(jsonResult))
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	jsonResult, _ := json.MarshalIndent(scheduleRecords[0], "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
