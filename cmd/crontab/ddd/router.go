package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/executeApi"
	"github.com/ryker-w/go-crontab/internal/api"
)

func Router(app *iris.Application) {
	root := app.Party("/api")
	admin(root.Party("/admin"))
	api.Router(root)
	return
}

func admin(root iris.Party)  {
	root.Post("/execute", executeApi.Execute)
}
