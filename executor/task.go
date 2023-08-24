package executor

import (
	"fmt"
	"github.com/lishimeng/go-log"
	common2 "github.com/ryker-w/go-crontab/common"
	"runtime/debug"
)

// FuncTask 任务执行函数
type FuncTask func(param common2.RunReq) (code int, msg string)

// Task 任务
type Task struct {
	Id        string
	Name      string
	Param     common2.RunReq
	fn        FuncTask
	StartTime int64
	EndTime   int64
}

// Run 运行任务
func (t *Task) Run() common2.CallElement {
	defer func() common2.CallElement {
		if err := recover(); err != nil {
			log.Info(err)
			debug.PrintStack() //堆栈跟踪
			return common2.Callback(t.Param, common2.Error, fmt.Sprintf("panic: %v", err))
		}
		return common2.Callback(t.Param, common2.Success, "")
	}()
	code, msg := t.fn(t.Param)
	return common2.Callback(t.Param, code, msg)
}

// Info 任务信息
func (t *Task) Info() string {
	return fmt.Sprintf("任务ID[%s]任务名称[%s]", t.Id, t.Name)
}
