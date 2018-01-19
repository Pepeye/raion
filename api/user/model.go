package user

import "gopkg.in/mgo.v2/bson"

// Schema Struct
type Schema struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	UUID   string        `json:"uuid" bson:"uuid"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
