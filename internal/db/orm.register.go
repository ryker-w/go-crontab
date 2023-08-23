package db

import (
	"github.com/ryker-w/go-crontab/internal/db/model"
)

func RegisterTables() (tables []interface{}) {
	tables = append(tables,
		new(model.OtterBatchTask),
		new(model.OtterBatchTaskLog),
		new(model.OtterBatchRunningTask),
	)
	return
}
