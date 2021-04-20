package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Schedule struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	HoVaTen    string             `bson:"HoVaTen" json:"fullName"`
	GhiChu     string             `bson:"GhiChu" json:"note"`
	LopKhoaHoc string             `bson:"LopKhoaHoc" json:"course"`
	MaSV       string             `bson:"MaSV" json:"studentId"`
	MaLMH      string             `bson:"MaLMH" json:"classId"`
	TenMonHoc  string             `bson:"TenMonHoc" json:"subject"`
	NgaySinh   string             `bson:"NgaySinh" json:"birthday"`
	Nhom       string             `bson:"Nhom" json:"group"`
	SoTinChi   int32              `bson:"SoTinChi" json:"credit"`
}

// SchedulesHandler returns some records from Schedule collection
func SchedulesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	CONNECTION_STRING := os.Getenv("CONNECTION_STRING")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		panic(err)
	}

	database := client.Database("TimetableUET")
	scheduleCollection := database.Collection("Schedule")
	filter := bson.M{"HoVaTen": "Ngô Quang Dương"}
	cursor, _ := scheduleCollection.Find(ctx, filter)

	scheduleRecords := make([]Schedule, 0)
	cursor.All(ctx, &scheduleRecords)

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	jsonResult, _ := json.MarshalIndent(scheduleRecords, "", "  ")
	fmt.Fprint(responseWriter, string(jsonResult))
}
