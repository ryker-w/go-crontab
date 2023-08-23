package model

import (
	"github.com/lishimeng/app-starter"
	"time"
)

// OtterBatchTask 定时任务管理
type OtterBatchTask struct {
	app.Pk
	ClientId        string    `orm:"column(client_id);null"`                // 任务应用、任务模块
	JobID           string    `orm:"column(job_id);unique"`                 // 任务ID。自动生成
	Name            string    `orm:"column(name);null"`                     // 任务名称
	ExecutorHandler string    `orm:"column(executor_handler)"`              // 任务标识。如“HelloWorld”
	Spec            string    `orm:"column(spec);null"`                     // cron表达式
	PrevTime        time.Time `orm:"column(prev_time);type(datetime);null"` // 上次执行时间
	NextTime        time.Time `orm:"column(next_time);type(datetime);null"` // 下次执行时间
	app.TableChangeInfo
}

const (
	Active   = 10 // 系统状态：激活
	Inactive = 20 // 系统状态：未激活
)

func (t *OtterBatchRunningTask) TableUnique() [][]string {
	return [][]string{
		{"ClientId", "ExecutorHandler"},
	}
}
