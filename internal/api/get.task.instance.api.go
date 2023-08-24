package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	common2 "github.com/ryker-w/go-crontab/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func getInstanceTaskApi(ctx iris.Context) {
	var resp common2.ScheduleResponse
	clientId := ctx.Params().Get("clientId")

	var rt model.OtterBatchRunningTask
	err := app.GetOrm().Context.QueryTable(new(model.OtterBatchRunningTask)).
		Filter("ClientId", clientId).
		Filter("Status", model.RunningStatusCreated).
		OrderBy("ctime").
		One(&rt)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	if rt.Id == 0 {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	rt.Status = model.RunningStatusExecuted
	_, err = app.GetOrm().Context.Update(&rt)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Data = common2.RunReq{
		JobID:           rt.JobId,
		ExecutorHandler: rt.ExecutorHandler,
		InstanceId:      rt.Id,
		Time:            rt.CreateTime,
		Data:            nil,
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
