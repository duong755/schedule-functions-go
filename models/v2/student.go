package modelsV2

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	StudentId       string             `bson:"studentId" json:"studentId,omitempty"`
	StudentName     string             `bson:"studentName" json:"studentName,omitempty"`
	StudentBirthday string             `bson:"studentBirthday" json:"studentBirthday,omitempty"`
	StudentCourse   string             `bson:"studentCourse" json:"studentCourse,omitempty"`
	ClassId         string             `bson:"classId" json:"classId,omitempty"`
	ClassNote       string             `bson:"classNote" json:"classNote,omitempty"`
	StudentNote     string             `bson:"studentNote" json:"studentNote,omitempty"`
	Classes         []Class            `json:"classes,omitempty"`
}
