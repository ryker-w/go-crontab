package util

import (
	"errors"
	"github.com/lishimeng/go-log"
	"github.com/robfig/cron/v3"
)

// ParseSpec 尝试转换定时任务表达式
func ParseSpec(spec string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Info("定时任务表达式(%+v)出错:%+v", spec, r)
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("发生未知panic")
			}
		}
	}()
	schedule, err := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
		).Parse(spec)
	if err != nil {
		return err
	}
	if schedule == nil {
		return errors.New("cron parse err")
	}
	return
}

// IsValidSpec 校验spec表达式是否合法
func IsValidSpec(spec string) (valid bool, err error) {
	err = ParseSpec(spec)
	if err == nil {
		log.Info("定时任务表达式(%+v)合法", spec)
		valid = true
	} else {
		log.Info("定时任务表达式(%+v)非法", spec)
		valid = false
	}
	return
}
