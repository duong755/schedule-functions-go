package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	Buoi       string             `bson:"Buoi" json:"session"`
	GhiChu     string             `bson:"GhiChu" json:"note"`
	GiangDuong string             `bson:"GiangDuong" json:"classroom"`
	GiaoVien   string             `bson:"GiaoVien" json:"lecturer"`
	MaLopMH    string             `bson:"MaLopMH" json:"classId"`
	MaMH       string             `bson:"MaMH" json:"subjectId"`
	SoSV       string             `bson:"SoSV" json:"numberOfStudents"`
	TenMonHoc  string             `bson:"TenMonHoc" json:"subjectName"`
	Thu        string             `bson:"Thu" json:"weekDay"`
	Tiet       string             `bson:"Tiet" json:"periods"`
	TinChi     string             `bson:"TinChi" json:"credits"`
}
