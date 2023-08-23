package model

import "github.com/lishimeng/app-starter"

// OtterBatchRunningTask 运行中任务
type OtterBatchRunningTask struct {
	app.Pk
	JobId           string `orm:"column(job_id)"`           // 任务ID
	ClientId        string `orm:"column(client_id)"`        // 任务应用、任务模块
	ExecutorHandler string `orm:"column(executor_handler)"` // 任务标识。如“HelloWorld”
	app.TableChangeInfo
}

const (
	RunningStatusCreated  = 0
	RunningStatusExecuted = 1
)
