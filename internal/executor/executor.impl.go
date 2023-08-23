package executor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/internal/common"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const defaultDelay = 1

type executor struct {
	ctx       context.Context
	host      string    // 调度服务地址
	clientId  string    // 执行器标识
	cancelReg bool      // 取消自动注册。默认false，即默认自动远程注册
	delay     int       // 自动获取延时。单位：秒。默认1s。
	regList   *ListTask // 注册任务列表
	runList   *ListTask // 正在执行任务列表
	mu        sync.RWMutex
}

func (e *executor) Run() (err error) {
	if len(e.host) == 0 || len(e.clientId) == 0 {
		return errors.New("host 和 clientId 不能为空")
	}
	if e.delay == 0 {
		e.delay = defaultDelay
	}
	if !e.cancelReg {
		err = e.RegTask()
		if err != nil {
			return
		}
	}
	e.run()
	return err
}

// AddRegTask 本地注册任务
func (e *executor) AddRegTask(handler string, task FuncTask) {
	if e.regList == nil {
		e.regList = &ListTask{
			data: make(map[string]*Task),
		}
	}
	var t = &Task{}
	t.fn = task
	e.regList.Set(handler, t)
	log.Info("执行器注册本地任务：%s", handler)
}

// RegTask 远程注册任务
func (e *executor) RegTask() (err error) {
	if len(e.clientId) == 0 {
		return errors.New("clientId 不能为空")
	}
	if e.regList == nil || e.runList.Len() == 0 {
		log.Info("无本地任务，取消远程注册")
		return
	}
	err = e.regTask()
	if err != nil {
		return err
	}
	return
}

func (e *executor) regTask() (err error) {
	req := common.RegReq{
		ClientId: e.clientId,
		Handlers: e.regList.GetKeys(),
	}
	var requestUrl = e.buildRemoteUrl(common.RouteRegTaskTpl)
	client := &http.Client{Timeout: common.HttpDefaultTimeout}
	jsonStr, _ := json.Marshal(req)
	resp, err := client.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Info("注册失败，client Post err")
		log.Info(err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	var response app.Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Info("注册失败，response Unmarshal err")
		log.Info(err)
		return
	}
	if response.Code != int64(tool.RespCodeSuccess) {
		return errors.New(fmt.Sprintf("注册失败, response code err: %d", response.Code))
	}
	return
}

func (e *executor) GetTaskInstance() common.RunReq {
	return e.getTaskInstance()
}

// RunTask 运行任务
func (e *executor) RunTask(req common.RunReq) common.CallElement {
	return e.runTask(req)
}

//运行一个任务
func (e *executor) runTask(req common.RunReq) common.CallElement {
	e.mu.Lock()
	defer e.mu.Unlock()
	// 本地执行
	if !e.regList.Exists(req.ExecutorHandler) {
		return common.Callback(req, common.Error, fmt.Sprintf("任务[ %d ]没有注册: %s", req.JobID, req.ExecutorHandler))
	}
	task := e.regList.Get(req.ExecutorHandler)
	task.Id = req.JobID
	task.Name = req.ExecutorHandler
	task.Param = req

	if e.runList == nil {
		e.runList = &ListTask{
			data: make(map[string]*Task),
		}
	}
	e.runList.Set(task.Id, task)
	log.Info("任务[%s]开始执行: %s", req.JobID, req.ExecutorHandler)
	return task.Run()
}

func (e *executor) getTaskInstance() common.RunReq {
	var requestUrl = e.buildRemoteUrl(common.RouteScheduleTpl)
	client := &http.Client{Timeout: common.HttpDefaultTimeout}
	resp, err := client.Get(requestUrl)
	if err != nil {
		log.Info("client Post err")
		log.Info(err)
		return common.RunReq{}
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	var response common.ScheduleResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Info("response Unmarshal err")
		log.Info(err)
		return common.RunReq{}
	}
	return response.Data
}

func (e *executor) buildRemoteUrl(routePath string) string {
	requestUrl, _ := url.JoinPath(e.host, fmt.Sprintf(routePath, e.clientId))
	return requestUrl
}

func (e *executor) response(callElement common.CallElement) error {
	var requestUrl = e.buildRemoteUrl(common.RouteScheduleTpl)
	client := &http.Client{Timeout: common.HttpDefaultTimeout}
	jsonStr, _ := json.Marshal(callElement)
	resp, err := client.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Info("client Post err")
		log.Info(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	result, _ := io.ReadAll(resp.Body)
	var response app.Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Info("response Unmarshal err")
		log.Info(err)
		return err
	}
	if response.Code != int64(tool.RespCodeSuccess) {
		return errors.New(fmt.Sprintf("response err code:%d", response.Code))
	}
	return nil
}

func (e *executor) run() {
	duration := time.Duration(e.delay) * time.Second
	go func() {
		var timer = time.NewTimer(duration)
		defer func() {
			timer.Stop()
		}()
		for {
			select {
			case <-e.ctx.Done():
				return
			case <-timer.C:
				req := e.GetTaskInstance()
				log.Info("GetTaskInstance: %+v", req)
				if req.InstanceId == 0 {
					callElement := common.Callback(req, common.NotFount, "")
					err := e.response(callElement)
					if err != nil {
						log.Info(err)
						return
					}
					return
				}
				callElement := e.RunTask(req)
				err := e.response(callElement)
				if err != nil {
					log.Info(err)
					return
				}
				timer.Reset(duration)
			}
		}
	}()
}