package tasklog

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func ListTaskLogApi(ctx iris.Context) {
	var resp app.PagerResponse
	clientId := ctx.URLParamDefault("clientId", "")
	name := ctx.URLParamDefault("name", "")
	jobId := ctx.URLParamDefault("jobId", "")
	executorHandler := ctx.URLParamDefault("executorHandler", "")
	result := ctx.URLParamIntDefault("result", 0)
	pageSize := ctx.URLParamIntDefault("pageSize", 10)
	pageNum := ctx.URLParamIntDefault("pageNum", 1)

	pager := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNum,
	}

	pager, taskLogs, err := getTaskLogList(pager, result, clientId, name, jobId, executorHandler)
	if err != nil {
		log.Info("[ListTaskLogApi]Get task log List err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	for _, taskLog := range taskLogs {
		pager.Data = append(pager.Data, common.TaskLog{
			Id:              taskLog.Id,
			ClientId:        taskLog.ClientId,
			JobId:           taskLog.JobId,
			Name:            taskLog.Name,
			ExecutorHandler: taskLog.ExecutorHandler,
			PrevTime:        tool.FormatTime(taskLog.PrevTime),
			NextTime:        tool.FormatTime(taskLog.NextTime),
			Result:          taskLog.Result,
			Msg:             taskLog.Msg,
			Ctime:           tool.FormatTime(taskLog.CreateTime),
		})
	}

	resp.Pager = pager
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func getTaskLogList(pager app.Pager, result int, clientId, name, jobId, executorHandler string) (p app.Pager, logs []model.OtterBatchTaskLog, err error) {
	//筛选项
	cond := orm.NewCondition()
	if result > 0 {
		cond = cond.And("result", result)
	}
	if len(clientId) > 0 {
		cond = cond.And("clientId", clientId)
	}
	if len(name) > 0 {
		cond = cond.And("name", name)
	}
	if len(jobId) > 0 {
		cond = cond.And("jobId", jobId)
	}
	if len(executorHandler) > 0 {
		cond = cond.And("executorHandler", executorHandler)
	}
	//sum
	var qs = app.GetOrm().Context.QueryTable(new(model.OtterBatchTaskLog)).SetCond(cond)
	sum, err := qs.Count()
	if err != nil {
		return
	}
	pager.TotalPage = common.CalcTotalPage(&pager, sum)

	//result_set
	_, err = qs.OrderBy("-ctime").Offset(common.CalcPageOffset(&pager)).Limit(pager.PageSize).All(&logs)
	if err != nil {
		return
	}
	p = pager
	return

}
