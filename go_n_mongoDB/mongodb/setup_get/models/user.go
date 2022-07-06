package models

import "gopkg.in/mgo.v2/bson"

//changing this
type User struct {
	Name   string        `json:"name"`
	Gender string        `json:"gender"`
	Age    int           `json:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}
