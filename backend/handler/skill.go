package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/highlanderdantas/gosquads/backend/lib/contx"
	"github.com/highlanderdantas/gosquads/backend/model"
	"github.com/highlanderdantas/gosquads/backend/repo"
)

//ListSkills lista skills
func ListSkills(ctx *contx.Context) {
	var query string
	if description := ctx.Query("description"); description != "" {
		query = description
	}
	skills, err := repo.ListSkill(query, ctx.Cache)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, skills)
}

//PageSkills lista skills paginados
func PageSkills(ctx *contx.Context) {
	var query string
	var pageable model.Pageable
	if page := ctx.Query("page"); page != "" {
		pageable.Page, _ = strconv.Atoi(page)
	}
	if size := ctx.Query("size"); size != "" {
		pageable.Size, _ = strconv.Atoi(size)
	}
	if description := ctx.Query("description"); description != "" {
		query = description
	}
	content, err := repo.PageSkill(query, pageable, ctx.Cache)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, content)
}

//GetSkill pega uma skill
func GetSkill(ctx *contx.Context) {
	id := ctx.Params("id")
	skill, err := repo.GetSkill(id, ctx.Cache)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, skill)
}

//DeleteSkill deleta uma skill
func DeleteSkill(ctx *contx.Context) {
	id := ctx.Params("id")
	err := repo.DeleteSkill(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, "Skill deletada com sucesso!")
}

//AddSkill adiciona um skill
func AddSkill(ctx *contx.Context) {
	body := ctx.GetBody(model.Skill{})
	err := repo.AddSkill(body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Skill adicionada com sucesso")
}
