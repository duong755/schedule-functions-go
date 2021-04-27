package modelsV2

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	Id               primitive.ObjectID `bson:"_id," json:"_id,omitempty"`
	SubjectId        string             `bson:"subjectId" json:"subjectId,omitempty"`
	SubjectName      string             `bson:"subjectName" json:"subjectName,omitempty"`
	Credit           int8               `bson:"credit" json:"credit,omitempty"`
	ClassId          string             `bson:"classId" json:"classId,omitempty"`
	Teacher          string             `bson:"teacher" json:"teacher,omitempty"`
	NumberOfStudents int8               `bson:"numberOfStudents" json:"numberOfStudents,omitempty"`
	Session          string             `bson:"session" json:"session,omitempty"` // "sáng" hoặc "chiều" hoặc "tối"
	WeekDay          int8               `bson:"weekDay" json:"weekDay,omitempty"`
	Periods          []int8             `bson:"periods" json:"periods,omitempty"`
	Place            string             `bson:"place" json:"place,omitempty"`
	Note             string             `bson:"note" json:"note,omitempty"` // "CL" (chung lớp), số (nhóm)
	Students         []Student          `bson:"students" json:"students,omitempty"`
	StudentNote      string             `bson:"studentNote" json:"studentNote,omitempty"`
	Groups           []Group            `bson:"groups" json:"groups,omitempty"`
}
