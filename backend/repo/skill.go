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

//ListSkill todas skills
func ListSkill(description string, cache cache.Cache) ([]model.Skill, error) {
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

//PageSkill todas skills paginadas
func PageSkill(description string, pageable model.Pageable, cache cache.Cache) (model.Content, error) {
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

//GetSkill pega uma skill
func GetSkill(id string, cache cache.Cache) (model.Skill, error) {
	skill := model.Skill{}
	query := "getskill:" + id

	if cache.IsExist(query) {
		skill = cache.Get(query).(model.Skill)
		return skill, nil
	}

	repository, err := conf.GetMongoCollection("skill")
	repository.FindId(bson.ObjectIdHex(id)).One(&skill)
	addCache(query, skill, cache)

	return skill, err
}

//DeleteSkill Deleta uma skill
func DeleteSkill(id string) error {
	repository, err := conf.GetMongoCollection("skill")
	err = repository.RemoveId(bson.ObjectIdHex(id))

	return err
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

//AddSkill teste
func AddSkill(body interface{}) error {
	repository, _ := conf.GetMongoCollection("skill")
	err := repository.Insert(body)
	return err
}

//addCache Adicionar skills em cache
func addCache(query string, content interface{}, cache cache.Cache) {
	s := fmt.Sprintf("%.0f", (time.Hour * 4).Seconds())
	timeout, _ := strconv.Atoi(s)
	cache.Put(query, content, int64(timeout))
}
