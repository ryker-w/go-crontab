package schedule

import (
	"github.com/lishimeng/go-log"
	"github.com/robfig/cron/v3"
	"github.com/ryker-w/go-crontab/internal/db/repo"
	"time"
)

type Job interface {
	Run() // 任务执行
}
type job struct {
	jobId    string
	spec     string
	clientId string
	entryID  cron.EntryID
	handler  string
	prevTime time.Time
	nextTime time.Time
	params   string
}

// Run 任务执行
func (j *job) run() {
	_, err := repo.AddRunningTask(j.jobId, j.clientId, j.handler, j.params)
	if err != nil {
		log.Info(err)
		log.Info("添加运行任务失败【%s】", j.handler)
		return
	}
}

func (j *job) Run() {
	j.run()
}
