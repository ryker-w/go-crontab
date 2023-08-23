package api

import (
	"github.com/kataras/iris/v12"
	"github.com/ryker-w/go-crontab/internal/common"
)

func Router(p iris.Party) {
	p.Get(common.RouteSchedulePath, getInstanceTaskApi)
	p.Post(common.RouteSchedulePath, responseInstanceTaskApi)

	p.Post(common.RouteRegTaskPath, RegisterTaskApi)
}
