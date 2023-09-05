package common

type Task struct {
	Id              int    `json:"id,omitempty"`
	ClientId        string `json:"clientId,omitempty"`
	JobId           string `json:"jobId,omitempty"`
	Name            string `json:"name,omitempty"`
	ExecutorHandler string `json:"executorHandler,omitempty"`
	Spec            string `json:"spec,omitempty"`
	PrevTime        string `json:"prevTime,omitempty"`
	NextTime        string `json:"nextTime,omitempty"`
	MaxInstanceNum  int    `json:"maxInstanceNum,omitempty"`
	InstanceNum     int    `json:"instanceNum,omitempty"`
	Status          int    `json:"status,omitempty"`
	Ctime           string `json:"ctime,omitempty"`
}

type TaskLog struct {
	Id              int    `json:"id,omitempty"`
	ClientId        string `json:"clientId,omitempty"`
	JobId           string `json:"jobId,omitempty"`
	Name            string `json:"name,omitempty"`
	ExecutorHandler string `json:"executorHandler,omitempty"`
	PrevTime        string `json:"prevTime,omitempty"`
	NextTime        string `json:"nextTime,omitempty"`
	Result          int    `json:"result,omitempty"`
	Msg             string `json:"msg,omitempty"`
	Ctime           string `json:"ctime,omitempty"`
}
