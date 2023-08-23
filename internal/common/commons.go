package common

import "time"

const RouteScheduleTpl = "/schedule/%s"
const RouteSchedulePath = "/schedule/{client}"

const RouteRegTaskTpl = "/register/%s"
const RouteRegTaskPath = "/register/{client}"

const HttpDefaultTimeout = 8 * time.Second

const (
	Success  = 200
	NotFount = 404
	Error    = 500
)
