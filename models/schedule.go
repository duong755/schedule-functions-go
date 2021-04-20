package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
