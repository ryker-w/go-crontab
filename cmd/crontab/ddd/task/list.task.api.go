package task

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func ListTaskApi(ctx iris.Context) {
	var resp app.PagerResponse
	status := ctx.Params().GetIntDefault("status", 0)
	pageSize := ctx.Params().GetIntDefault("pageSize", 10)
	pageNum := ctx.Params().GetIntDefault("pageNum", 1)

	pager := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNum,
	}

	pager, tasks, err := getTaskList(pager, status)
	if err != nil {
		log.Info("[ListTaskApi]Get task List err")
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	}

	for _, task := range tasks {
		var instanceNum int64
		instanceNum, err = app.GetOrm().Context.QueryTable(new(model.OtterBatchRunningTask)).Filter("client_id", task.ClientId).Filter("job_id", task.JobID).Count()
		if err != nil {
			log.Info("[ListTaskApi]Count task err, id=%d", task.Id)
			log.Info(err)
			resp.Code = tool.RespCodeError
			resp.Message = err.Error()
			tool.ResponseJSON(ctx, resp)
			return
		}
		pager.Data = append(pager.Data, common.Task{
			Id:              task.Id,
			ClientId:        task.ClientId,
			JobId:           task.JobID,
			Name:            task.Name,
			ExecutorHandler: task.ExecutorHandler,
			Spec:            task.Spec,
			PrevTime:        tool.FormatTime(task.PrevTime),
			NextTime:        tool.FormatTime(task.NextTime),
			MaxInstanceNum:  task.MaxInstanceNum,
			InstanceNum:     int(instanceNum),
			Status:          task.Status,
			Ctime:           tool.FormatTime(task.CreateTime),
		})
	}

	resp.Pager = pager
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func getTaskList(pager app.Pager, status int) (p app.Pager, tasks []model.OtterBatchTask, err error) {
	//筛选项
	cond := orm.NewCondition()
	if status > 0 {
		cond = cond.And("status", status)
	}
	//sum
	var qs = app.GetOrm().Context.QueryTable(new(model.OtterBatchTask)).SetCond(cond)
	sum, err := qs.Count()
	if err != nil {
		return
	}
	pager.TotalPage = common.CalcTotalPage(&pager, sum)

	//result_set
	_, err = qs.OrderBy("Id").Offset(common.CalcPageOffset(&pager)).Limit(pager.PageSize).SetCond(cond).All(&tasks)
	if err != nil {
		return
	}
	p = pager
	return

}
