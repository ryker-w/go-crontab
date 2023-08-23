package executor

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/internal/common"
)

// Router 远程服务器接口
func Router(p iris.Party) {
	p.Post(common.RouteSchedulePath, apiExecutor)
}

func apiExecutor(ctx iris.Context) {
	log.Info("Execute")
	var req common.RunReq
	var resp common.Response

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = "请求参数错误！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	callElement := GetExecutor().RunTask(common.RunReq{
		JobID:           req.JobID,
		ExecutorHandler: req.ExecutorHandler,
		Data:            nil,
	})
	resp.Code = tool.RespCodeSuccess
	resp.Data = callElement
	tool.ResponseJSON(ctx, resp)
}
