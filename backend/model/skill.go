package model

import "gopkg.in/mgo.v2/bson"

//Skill habilidade de um usuario
type Skill struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Description string        `bson:"description" json:"description"`
}
