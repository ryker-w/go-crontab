package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/task"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/tasklog"
	"github.com/ryker-w/go-crontab/internal/api"
)

func Router(app *iris.Application) {
	root := app.Party("/api")
	admin(root.Party("/admin"))
	api.Router(root)
	return
}

func admin(root iris.Party) {
	root.Get("/task", task.ListTaskApi)
	root.Post("/task/spec", task.UpdateTaskSpecApi)
	root.Post("/task/start", task.StartTaskApi)
	root.Post("/task/kill", task.KillTaskApi)
	root.Post("/task/exec", task.ExecuteTaskApi)
	root.Put("/task", task.UpdateTaskInfoApi)

	root.Get("/tasklog", tasklog.ListTaskLogApi)
}
