package task

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/internal/db/model"
	"github.com/ryker-w/go-crontab/internal/db/repo"
	"github.com/ryker-w/go-crontab/internal/schedule"
)

type executeReq struct {
	Id     int         `json:"id,omitempty"`
	Params interface{} `json:"params,omitempty"`
}

func ExecuteTaskApi(ctx iris.Context) {
	log.Info("Execute")
	var req executeReq
	var resp app.Response

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("[ExecuteTaskApi]ReadJSON err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	var task model.OtterBatchTask
	task.Id = req.Id
	err = app.GetOrm().Context.Read(&task)
	if err != nil {
		log.Info("[ExecuteTaskApi]Get task err, id=%d", req.Id)
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}
	if task.Status == model.Inactive {
		log.Info("[ExecuteTaskApi] Task inactive, id=%d", req.Id)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if !schedule.Sch.ExistJob(task.JobID) {
		log.Info("[ExecuteTaskApi] schedule Task not exist, id=%d", req.Id)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	var params = make([]byte, 0)
	if req.Params != nil {
		params, err = json.Marshal(req.Params)
		if err != nil {
			log.Info("[ExecuteTaskApi] Marshal err, Params=%+v", req.Params)
			resp.Code = tool.RespCodeError
			resp.Message = "任务参数错误"
			tool.ResponseJSON(ctx, resp)
			return
		}
	}

	_, err = repo.AddRunningTask(task.JobID, task.ClientId, task.ExecutorHandler, string(params))
	if err != nil {
		log.Info(err)
		log.Info("[ExecuteTaskApi] 添加运行任务失败, Id=%d", req.Id)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
