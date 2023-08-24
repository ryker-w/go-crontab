package executor

import (
	common2 "github.com/ryker-w/go-crontab/common"
)

// Executor 执行器
type Executor interface {
	Run() (err error)
	// AddRegTask 本地注册
	AddRegTask(handler string, task FuncTask)
	// RegTask 服务器注册
	RegTask() (err error)
	// GetTaskInstance 获取调度任务实例
	GetTaskInstance() common2.RunReq
	// RunTask 运行任务
	RunTask(req common2.RunReq) common2.CallElement
}
