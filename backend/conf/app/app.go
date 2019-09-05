package app

import (
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"github.com/highlanderdantas/gosquads/backend/conf"
	"github.com/highlanderdantas/gosquads/backend/handler"
	"github.com/highlanderdantas/gosquads/backend/lib/cache"
	"github.com/highlanderdantas/gosquads/backend/lib/contx"
	"github.com/highlanderdantas/gosquads/backend/lib/cors"
	"github.com/highlanderdantas/gosquads/backend/lib/template"
	"gopkg.in/macaron.v1"
)

//SetupMiddlewares configures the middlewares using in each web request
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))
	app.Use(macaron.Static("public"))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	//Cache in memory
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	app.Use(session.Sessioner())
	app.Use(contx.Contexter())
	app.Use(cors.Cors())

}

//SetupRoutes defines the routes the Web Application will respond
func SetupRoutes(app *macaron.Macaron) {
	app.Get("", func() string {
		return "Lets gosquads!"
	})
	app.Group("/skill", func() {
		app.Get("", handler.ListSkills)
		app.Get("/:id", handler.GetSkill)
		app.Post("", handler.AddSkill)
		app.Get("/page", handler.PageSkills)
		app.Delete("/:id", handler.DeleteSkill)
	})

}
