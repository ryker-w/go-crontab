package test

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/ryker-w/go-crontab/common"
	"github.com/ryker-w/go-crontab/executor"
	"sync"
	"testing"
)

func TestClient(t *testing.T) {
	var w sync.WaitGroup
	w.Add(1)

	var ctx = context.Background()
	err := createExecutorClient(ctx)
	if err != nil {
		t.Fatal(err)
		return
	}

	w.Wait()
	t.Log("test finish")
}

func createExecutorClient(ctx context.Context) error {
	err := executor.NewExecutor(ctx,
		executor.WithHost("https://open.thingplecloud.com/crontab"),
		executor.WithClientId("test_random_client_qwsed5rf6tg7h8"),
		executor.WithTask("demo1_hello_world", demo1HelloWorld),
		executor.WithTask("demo2_task", demo2Task),
		executor.WithTask("demo3_task", func(param common.RunReq) (code int, msg string) {
			log.Info("[demo3_task] start")
			return common.Error, "error"
		}),
	)
	if err != nil {
		return err
	}
	return nil
}



func demo1HelloWorld(req common.RunReq) (code int, msg string)  {
	log.Info("[helloWorld] running time: %v", req.Time)

	return common.Success, ""
}

func demo2Task(req common.RunReq) (code int, msg string)  {
	log.Info("[demoTask1] receiving data: %+v", req.Data)

	return common.Success, ""
}
