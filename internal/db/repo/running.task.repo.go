package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func AddRunningTask(jobId string, clientId string, handler string) (instanceId int, err error) {
	var task model.OtterBatchRunningTask
	task.JobId = jobId
	task.ClientId = clientId
	task.ExecutorHandler = handler
	task.Status = model.RunningStatusCreated
	id, err := app.GetOrm().Context.Insert(&task)
	instanceId = int(id)
	return
}

func DeleteRunningTask(id int) (err error) {
	var rt model.OtterBatchRunningTask
	rt.Id = id
	_, err = app.GetOrm().Context.Delete(&rt)
	return
}
