package task

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func UpdateTaskInfoApi(ctx iris.Context) {
	var req common.Task
	var resp app.Response

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("[UpdateTaskInfoApi]ReadJSON err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	if req.Id == 0 {
		log.Info("[UpdateTaskInfoApi]Id=%d || spec=%s, nil", req.Id, req.Spec)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	var task model.OtterBatchTask
	task.Id = req.Id
	err = app.GetOrm().Context.Read(&task)
	if err != nil {
		log.Info("[UpdateTaskInfoApi]Get task err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	task.Name = req.Name
	task.MaxInstanceNum = req.MaxInstanceNum
	_, err = app.GetOrm().Context.Update(&task, "Name", "MaxInstanceNum")
	if err != nil {
		log.Info("[UpdateTaskInfoApi]Update task err, id=%s", req.Id)
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

