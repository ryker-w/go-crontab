package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	common2 "github.com/ryker-w/go-crontab/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
	"time"
)

func GetActiveTasks() (sum int, tasks []model.OtterBatchTask, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.OtterBatchTask)).Filter("status", model.Active)
	count, err := qs.Count()
	if err != nil {
		return
	}
	sum = int(count)
	_, err = qs.All(&tasks)
	if err != nil {
		return
	}
	return
}

func HandleTaskResult(callElement common2.CallElement, nextTime time.Time) (err error) {
	// TODO 添加任务记录
	task, err := FindTaskByJobId(callElement.JobId)
	if err != nil {
		return
	}
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		task.PrevTime = task.NextTime
		task.NextTime = nextTime
		_, e = ctx.Context.Update(&task, "NextTime", "PrevTime")
		if e != nil {
			return e
		}
		// log
		l := model.OtterBatchTaskLog{
			ClientId:        task.ClientId,
			JobId:           task.JobID,
			Name:            task.Name,
			ExecutorHandler: task.ExecutorHandler,
			PrevTime:        task.PrevTime,
			NextTime:        task.NextTime,
			Result:          callElement.HandleCode,
			Msg:             callElement.HandleMsg,
		}
		_, e = ctx.Context.Insert(&l)
		if e != nil {
			return e
		}
		err = DeleteRunningTask(callElement.InstanceId)
		if err != nil {
			log.Info("运行中任务删除失败【%s】", callElement.InstanceId)
			return
		}
		return e
	})
	if err != nil {
		return err
	}

	return
}

func FindTaskByJobId(jobId string) (task model.OtterBatchTask, err error) {
	task.JobID = jobId
	err = app.GetOrm().Context.Read(&task, "JobID")
	return
}


func CancelTaskByJobId(jobId string) (err error) {
	var task model.OtterBatchTask
	task.JobID = jobId
	err = app.GetOrm().Context.Read(&task, "JobID")
	if err != nil {
		return
	}
	task.Status = model.Inactive
	_, err = app.GetOrm().Context.Update(&task, "status")
	if err != nil {
		return
	}
	return
}
