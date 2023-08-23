package setup

import (
	"context"
	"github.com/ryker-w/go-crontab/internal/schedule"
)

func Setup(ctx context.Context) (err error) {
	// 初始化调度器
	schedule.New()
	// 测试

	return
}
