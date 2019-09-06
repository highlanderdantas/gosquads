package repo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/highlanderdantas/gosquads/backend/conf"
	"github.com/highlanderdantas/gosquads/backend/model"
)

//AddUser adiciona um usuario
func AddUser(user model.UserDTO) error {
	repository, _ := conf.GetMongoCollection("user")
	skills := []mgo.DBRef{}

	for _, skill := range user.Skills {
		skills = append(skills, mgo.DBRef{Collection: "skill", Id: bson.ObjectIdHex(skill.ID)})
	}

	id := bson.NewObjectId()
	err := repository.Insert(&model.User{ID: id, Name: user.Name, Squad: user.Squad, Head: user.Head, Skills: skills})
	return err
}

//ListUsers todas users
func ListUsers() ([]model.User, error) {
	users := []model.User{}
	repository, err := conf.GetMongoCollection("user")
	repository.Find(nil).All(&users)
	return users, err
}
