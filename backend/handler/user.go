package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/highlanderdantas/gosquads/backend/lib/contx"
	"github.com/highlanderdantas/gosquads/backend/model"
	"github.com/highlanderdantas/gosquads/backend/repo"
)

//AddUser adiciona um user
func AddUser(ctx *contx.Context) {
	body, err := getUser(ctx)
	err = repo.AddUser(body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Usuario adicionado com sucesso")
}

//ListUsers lista users
func ListUsers(ctx *contx.Context) {
	skills, err := repo.ListUsers()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, skills)
}

func getUser(ctx *contx.Context) (model model.UserDTO, err error) {
	body, err := ctx.Req.Body().Bytes()
	defer ctx.Req.Body().ReadCloser()
	err = json.Unmarshal(body, &model)

	return model, err
}
