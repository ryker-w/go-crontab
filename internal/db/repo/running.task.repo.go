package repo

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func AddRunningTask(jobId, clientId, handler, params string) (instanceId int, err error) {
	//
	var task model.OtterBatchTask
	task.JobID = jobId
	err = app.GetOrm().Context.Read(&task, "job_id")
	if err != nil {
		return
	}
	if task.MaxInstanceNum != 0 {
		count, e := app.GetOrm().Context.QueryTable(new(model.OtterBatchRunningTask)).
			Filter("executor_handler", handler).Filter("client_id", clientId).
			Count()
		if e != nil {
			return instanceId, e
		}
		sum := int(count)
		if sum == task.MaxInstanceNum || sum > task.MaxInstanceNum {
			return instanceId, errors.New("超出最大可执行限制")
		}
	}
	var rt model.OtterBatchRunningTask
	rt.JobId = jobId
	rt.ClientId = clientId
	rt.ExecutorHandler = handler
	rt.Status = model.RunningStatusCreated
	rt.Params = params
	id, err := app.GetOrm().Context.Insert(&rt)
	instanceId = int(id)
	return
}

func DeleteRunningTask(id int) (err error) {
	var rt model.OtterBatchRunningTask
	rt.Id = id
	_, err = app.GetOrm().Context.Delete(&rt)
	return
}
