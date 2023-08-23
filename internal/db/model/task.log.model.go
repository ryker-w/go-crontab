package model

import (
	"github.com/lishimeng/app-starter"
	"time"
)

// OtterBatchTaskLog 任务日志
type OtterBatchTaskLog struct {
	app.Pk
	ClientId        string    `orm:"column(client_id);"`                    // 任务应用、任务模块
	JobId           string    `orm:"column(job_id);null"`                   // 任务ID
	Name            string    `orm:"column(name);null"`                     // 任务名称
	ExecutorHandler string    `orm:"column(executor_handler);null"`         // 任务标识。如“HelloWorld”
	PrevTime        time.Time `orm:"column(prev_time);type(datetime);null"` // 上次执行时间
	NextTime        time.Time `orm:"column(next_time);type(datetime);null"` // 下次执行时间
	Result          int       `orm:"column(result);null"`                   // 执行结果
	Msg             string    `orm:"column(msg);null"`                      // 错误信息
	app.TableInfo
}
