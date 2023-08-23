package main

import (
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/ryker-w/go-crontab/cmd"
	"github.com/ryker-w/go-crontab/cmd/crontab/ddd"
	"github.com/ryker-w/go-crontab/internal/db"
	"github.com/ryker-w/go-crontab/internal/etc"
	"github.com/ryker-w/go-crontab/setup"
	"time"
)
import _ "github.com/lib/pq"

func main() {
	orm.Debug = true

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(cmd.AppName)
	fmt.Println(cmd.Version)

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		builder.EnableDatabase(dbConfig.Build(),
			db.RegisterTables()...).
			SetWebLogLevel("debug").
			EnableWeb(etc.Config.Web.Listen, ddd.Router).
			ComponentBefore(setup.Setup)
		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
