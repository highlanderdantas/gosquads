package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserDTO dto
type UserDTO struct {
	ID     string     `bson:"_id" json:"id"`
	Name   string     `bson:"name" json:"name"`
	Head   bool       `bson:"head" json:"head"`
	Squad  string     `bson:"squad" json:"squad"`
	Skills []SkillDTO `bson:"skills" json:"skills"`
}

//User entiny
type User struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Head   bool          `bson:"head" json:"head"`
	Squad  string        `bson:"squad" json:"squad"`
	Skills []mgo.DBRef   `bson:"skills" json:"skills"`
}
