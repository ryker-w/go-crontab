package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/ryker-w/go-crontab/internal/common"
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func RegisterTaskApi(ctx iris.Context) {
	var req common.RegReq
	var resp app.Response
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.ClientId) == 0 || len(req.Handlers) == 0 {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	err = app.GetOrm().Transaction(func(tx persistence.TxContext) (e error) {
		for _, handler := range req.Handlers {
			var count int64
			count, e = tx.Context.QueryTable(new(model.OtterBatchTask)).
				Filter("ClientId", req.ClientId).
				Filter("ExecutorHandler", handler).Count()
			if e != nil {
				return e
			}
			if count != 0 {
				continue
			}
			t := model.OtterBatchTask{
				ClientId:        req.ClientId,
				JobID:           tool.GetRandomString(8),
				Name:            handler,
				ExecutorHandler: handler,
				Spec:            "",
				TableChangeInfo: app.TableChangeInfo{
					Status: model.Inactive,
				},
			}
			_, e = tx.Context.Insert(&t)
			if e != nil {
				return e
			}
		}
		return
	})
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
