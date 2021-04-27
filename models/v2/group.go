package modelsV2

type Group struct {
	Session string `bson:"session" json:"session,omitempty"`
	WeekDay int32  `bson:"weekDay" json:"weekDay,omitempty"`
	Place   string `bson:"place" json:"place,omitempty"`
	Teacher string `bson:"teacher" json:"teacher,omitempty"`
	Note    string `bson:"note" json:"note,omitempty"`
}
