package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/internal/common"
	"github.com/ryker-w/go-crontab/internal/schedule"
)

func responseInstanceTaskApi(ctx iris.Context) {
	var req common.CallElement
	var resp app.Response
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	if req.HandleCode == common.Success {
		log.Info("任务执行成功【%s】", req.ExecutorHandler)
		schedule.Sch.CallbackJob(req)
	} else if req.HandleCode == common.NotFount {
		log.Info("任务未执行【%s】", req.ExecutorHandler)
		schedule.Sch.CallbackJob(req)
	} else {
		log.Info("任务执行失败【%s】, code: %d, msg:%s", req.ExecutorHandler, req.HandleCode, req.HandleMsg)
		schedule.Sch.CallbackJob(req)
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
