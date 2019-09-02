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
	return skills, err
}

//addCache Adicionar skills em cache
func addCache(query string, skills []model.Skill, cache cache.Cache) {
	s := fmt.Sprintf("%.0f", (time.Hour * 4).Seconds())
	timeout, _ := strconv.Atoi(s)
	cache.Put(query, skills, int64(timeout))
}
