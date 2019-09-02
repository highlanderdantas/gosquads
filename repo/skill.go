package repo

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-macaron/cache"
	"github.com/highlanderdantas/gosquads/conf"
	"github.com/highlanderdantas/gosquads/model"
)

//FindAll todas skills
func FindAll(description string, cache cache.Cache) ([]model.Skill, error) {
	skills := []model.Skill{}
	if cache.IsExist(description) {
		skills = cache.Get(description).([]model.Skill)
		return skills, nil
	}
	repository, err := conf.GetMongoCollection("skill")
	if description == "" {
		repository.Find(nil).All(&skills)
	} else {
		repository.Find(bson.M{"description": bson.RegEx{Pattern: description, Options: "i"}}).All(&skills)
	}
	addCache(description, skills, cache)
	return skills, err
}

//FindAllPageable todas skills paginadas
func FindAllPageable(description string, pageable model.Pageable, cache cache.Cache) ([]model.Skill, error) {
	skills := []model.Skill{}
	repository, err := conf.GetMongoCollection("skill")
	if description == "" {
		repository.Find(nil).Skip(pageable.Page).Limit(pageable.Size).All(&skills)
	} else {
		repository.Find(bson.M{"description": bson.RegEx{Pattern: description, Options: "i"}}).Skip(pageable.Page).Limit(pageable.Size).All(&skills)
	}
	return skills, err
}

//addCache Adicionar skills em cache
func addCache(query string, skills []model.Skill, cache cache.Cache) {
	s := fmt.Sprintf("%.0f", (time.Hour * 4).Seconds())
	timeout, _ := strconv.Atoi(s)
	cache.Put(query, skills, int64(timeout))
}
