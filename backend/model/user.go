package model

import "gopkg.in/mgo.v2/bson"

//User entiny
type User struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Head   bool          `bson:"head" json:"head"`
	Squad  string        `bson:"squad" json:"squad"`
	Skills []Skill       `bson:"skills" json:"skills"`
}
