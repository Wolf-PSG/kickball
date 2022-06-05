package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Players struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Name               string             `bson:"name"`
	Speed_attribute    int32              `bson:"speed_attribute"`
	Power_attribute    int32              `bson:"power_attribute"`
	Accuracy_attribute int32              `bson:"accuracy_attribute"`
	Defence_attribute  int32              `bson:"defence_attribute"`
	Passing_attribute  int32              `bson:"passing_attribute"`
	Style              string             `bson:"style"`
	Corner_preference  string             `bson:"corner_preference"`
	Skill              string             `bson:"skill"`
	Division           string             `bson:"division"`
}

var PlayerAttributes = []string{
	"speed",
	"power",
	"accuracy",
	"defence",
	"passing",
}
