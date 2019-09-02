package repo

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-macaron/cache"
	"github.com/highlanderdantas/gosquads/backend/conf"
	"github.com/highlanderdantas/gosquads/backend/model"
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
func FindAllPageable(description string, pageable model.Pageable, cache cache.Cache) (model.Content, error) {
	skills := []model.Skill{}
	repository, err := conf.GetMongoCollection("skill")
	if description == "" {
		repository.Find(nil).Skip(pageable.Page).Limit(pageable.Size).All(&skills)
	} else {
		repository.Find(bson.M{"description": bson.RegEx{Pattern: description, Options: "i"}}).Skip(pageable.Page).Limit(pageable.Size).All(&skills)
	}

	content := model.Content{
		Content:       skills,
		TotalElements: Count(cache),
		Pageable:      pageable,
	}
	return content, err
}

//Count traz a quantidade de registros
func Count(cache cache.Cache) int {
	query := "cacheCountSkills"
	if cache.IsExist(query) {
		return cache.Get(query).(int)
	}
	repository, _ := conf.GetMongoCollection("skill")
	totalElements, _ := repository.Find(nil).Count()
	addCache(query, totalElements, cache)

	return totalElements
}

//addCache Adicionar skills em cache
func addCache(query string, content interface{}, cache cache.Cache) {
	s := fmt.Sprintf("%.0f", (time.Hour * 4).Seconds())
	timeout, _ := strconv.Atoi(s)
	cache.Put(query, content, int64(timeout))
}
