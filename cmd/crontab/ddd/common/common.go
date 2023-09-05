package common

import (
	"github.com/lishimeng/app-starter"
	"math"
)

const (
	RespCodeSuccess       = 200
	RespCodeNotFound      = 404
	RespCodeNeedAuth      = 401
	RespCodeNeedOrg       = 402
	RespCodeInternalError = 500
)

func CalcTotalPage(p *app.Pager, count int64) int {
	t := math.Ceil(float64(count) / float64(p.PageSize))
	return int(t)
}

func CalcPageOffset(p *app.Pager) int {
	return (p.PageNum - 1) * p.PageSize
}
