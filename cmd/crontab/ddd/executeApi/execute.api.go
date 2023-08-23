package executeApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type ExecuteReq struct {
	JobID           string `json:"jobID"`
	ExecutorHandler string `json:"executorHandler"`
}

func Execute(ctx iris.Context) {
	log.Info("Execute")
	var req ExecuteReq
	var resp app.Response

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	//callElement := executor.GetExecutor().RunTask(common.RunReq{
	//	JobID:           req.JobID,
	//	ExecutorHandler: req.ExecutorHandler,
	//	Data:            nil,
	//})
	//
	//if callElement.HandleCode != common.Success {
	//	resp.Code = tool.RespCodeError
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
