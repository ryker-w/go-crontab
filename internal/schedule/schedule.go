package schedule

import (
	"github.com/lishimeng/go-log"
	"github.com/robfig/cron/v3"
	"github.com/ryker-w/go-crontab/internal/common"
	"github.com/ryker-w/go-crontab/internal/db/repo"
)

type Schedule interface {
	Init()                                                                   // 初始化定时任务
	AddJob(jobId string, executorHandler string, spec string, remote string) // 动态添加定时任务
	DelJob(jobId string)                                                     // 动态删除定时任务
	CallbackJob(element common.CallElement)                                  // 定时任务结果回调
}

type schedule struct {
	cron    *cron.Cron
	jobList *jobList
}

func (s *schedule) Init() {
	s.cron = cron.New(cron.WithSeconds())
	s.cron.Start()
	s.jobList = &jobList{
		data: make(map[string]*job),
	}

	sum, tasks, err := repo.GetActiveTasks()
	if err != nil {
		log.Info("初始化任务失败： %v", err)
		return
	}
	log.Info("调度器初始化定时任务...共 %d 个", sum)
	for _, task := range tasks {
		s.addJob(task.JobID, task.ExecutorHandler, task.Spec, task.ClientId)
	}
}

func (s *schedule) addJob(jobId string, executorHandler string, spec string, clientId string) {
	j := job{
		jobId:    jobId,
		spec:     spec,
		handler:  executorHandler,
		clientId: clientId,
	}
	exists := s.jobList.Exists(jobId)
	if exists {
		log.Info("任务[ %s ]已存在，添加失败。handler: %s", jobId, executorHandler)
		return
	}
	entryID, err := s.cron.AddJob(spec, &j)
	if err != nil {
		log.Info(err)
		log.Info("任务[ %s ]定时格式错误，添加失败。handler: %s", jobId, executorHandler)
		return
	}
	entry := s.cron.Entry(entryID)
	log.Info("任务初始化[成功]jobId:%s, entryID: %d, handler: %s, 下次执行时间：%s", jobId, entryID, executorHandler, entry.Next.Format("2006-01-02 15:04:05"))
	j.entryID = entryID
	j.nextTime = entry.Next
	s.jobList.Set(jobId, &j)
}

func (s *schedule) callbackJob(callElement common.CallElement) {
	j := s.jobList.Get(callElement.JobId)
	entry := s.cron.Entry(j.entryID)
	j.nextTime = entry.Next
	s.jobList.Set(callElement.JobId, j)
	//TODO 更新数据库任务：下次执行时间
	err := repo.HandleTaskResult(callElement, entry.Next)
	if err != nil {
		log.Info("更新任务下次执行时间失败[%d]%s, err: %v", j.entryID, j.handler, err)
		return
	}
	// TODO 无执行函数，取消任务
	if callElement.HandleCode == common.NotFount {
		s.delJob(callElement.JobId)
		err = repo.CancelTaskByJobId(callElement.JobId)
		if err != nil {
			log.Info("取消任务失败[%d]%s, err: %v", j.entryID, j.handler, err)
			return
		}
	}
}

func (s *schedule) delJob(jobId string) {
	j := s.jobList.Get(jobId)
	log.Info("[schedule]删除定时任务: %s, %s", j.jobId, j.handler)
	s.cron.Remove(j.entryID)
	s.jobList.Del(jobId)
}

func (s *schedule) AddJob(jobId string, executorHandler string, spec string, remote string) {
	s.addJob(jobId, executorHandler, spec, remote)
}

func (s *schedule) CallbackJob(callElement common.CallElement) () {
	s.callbackJob(callElement)
}

func (s *schedule) DelJob(jobId string) {
	s.delJob(jobId)
}
