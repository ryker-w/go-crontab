package task

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
	"github.com/ryker-w/go-crontab/internal/schedule"
	"github.com/ryker-w/go-crontab/internal/util"
)

func UpdateTaskSpecApi(ctx iris.Context) {
	var req common.Task
	var resp app.Response

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("[UpdateTaskSpecApi]ReadJSON err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	if req.Id == 0 || len(req.Spec) == 0 {
		log.Info("[UpdateTaskSpecApi]Id=%d || spec=%s, nil", req.Id, req.Spec)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	var task model.OtterBatchTask
	task.Id = req.Id
	err = app.GetOrm().Context.Read(&task)
	if err != nil {
		log.Info("[UpdateTaskSpecApi]Get task err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	//TODO 检查cron格式
	valid, err := util.IsValidSpec(req.Spec)
	if err != nil {
		log.Info("[UpdateTaskSpecApi]IsValidSpec err, spec=%s, valid=%v", req.Spec, valid)
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	if !valid {
		log.Info("[UpdateTaskSpecApi]spec=%s, valid=%v", req.Spec, valid)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	task.Spec = req.Spec
	task.Status = model.Inactive
	_, err = app.GetOrm().Context.Update(&task, "spec", "status")
	if err != nil {
		log.Info("[UpdateTaskSpecApi]Update task err, id=%s", req.Id)
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	//终止任务
	schedule.Sch.DelJob(task.JobID)

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
