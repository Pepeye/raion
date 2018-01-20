package user

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const resourceName string = "users"

// Schema Struct
type Schema struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	UUID   string        `json:"uuid" bson:"uuid"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

// setup indexes
func init() {
	db, err := mgo.Dial("mongodb://localhost")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	// fetch collection
	col := db.DB("raion").C(resourceName)

	index := mgo.Index{
		Key:        []string{"uuid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = col.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[raion]: create database indexes for resource - %s\n", resourceName)
}
