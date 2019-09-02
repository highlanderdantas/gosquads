package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/highlanderdantas/gosquads/lib/contx"
	"github.com/highlanderdantas/gosquads/repo"
)

//ListSkills lista skills
func ListSkills(ctx *contx.Context) {
	var query string
	if description := ctx.Query("description"); description != "" {
		query = description
	}
	fmt.Println(query)
	skills, err := repo.FindAll(query, ctx.Cache)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, skills)
}
