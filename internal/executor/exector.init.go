package executor

import (
	"context"
)

var exec Executor

// NewExecutor 创建执行器
func NewExecutor(ctx context.Context, options ...Option) {
	e := &executor{
		ctx: ctx,
		regList: &ListTask{
			data: make(map[string]*Task),
		},
		runList: &ListTask{
			data: make(map[string]*Task),
		},
	}
	for _, o := range options {
		o(e)
	}
	err := e.Run()
	if err != nil {
		return
	}
}

type Option func(e *executor)

func WithHost(host string) Option {
	return func(e *executor) {
		e.host = host
		//todo 测试连接
	}
}

func WithClientId(clientId string) Option {
	return func(e *executor) {
		e.clientId = clientId
	}
}

// WithTask 本地注册任务
func WithTask(handler string, task FuncTask) Option {
	return func(e *executor) {
		e.AddRegTask(handler, task)
	}
}

func WithDelay(delay int) Option {
	return func(e *executor) {
		e.delay = delay
	}
}

func WithCancelReg(cancel bool) Option {
	return func(e *executor) {
		e.cancelReg = cancel
	}
}

func GetExecutor() Executor {
	return exec
}
