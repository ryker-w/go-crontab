package task

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
	"github.com/ryker-w/go-crontab/internal/schedule"
)

func KillTaskApi(ctx iris.Context)  {
	var req common.Task
	var resp app.Response

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("[KillTaskApi]ReadJSON err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	if req.Id == 0 {
		log.Info("[KillTaskApi]Id nil")
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	var task model.OtterBatchTask
	task.Id = req.Id
	err = app.GetOrm().Context.Read(&task)
	if err != nil {
		log.Info("[KillTaskApi]Get task err, id=%d", req.Id)
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	task.Status = model.Inactive
	_, err = app.GetOrm().Context.Update(&task, "status")
	if err != nil {
		log.Info("[KillTaskApi]Update task err, id=%d", req.Id)
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	if schedule.Sch.ExistJob(task.JobID) {
		schedule.Sch.DelJob(task.JobID)
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

