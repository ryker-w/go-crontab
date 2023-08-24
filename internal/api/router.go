package api

import (
	"github.com/kataras/iris/v12"
	common2 "github.com/ryker-w/go-crontab/common"
)

func Router(p iris.Party) {
	p.Get(common2.RouteSchedulePath, getInstanceTaskApi)
	p.Post(common2.RouteSchedulePath, responseInstanceTaskApi)

	p.Post(common2.RouteRegTaskPath, RegisterTaskApi)
}
