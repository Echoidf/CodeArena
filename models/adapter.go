package models

import (
	"CodeArena/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Engine *xorm.Engine

func init() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.V.GetString("mysql.user"),
		conf.V.GetString("mysql.pwd"),
		conf.V.GetString("mysql.host"),
		conf.V.GetInt("mysql.port"),
		conf.V.GetString("mysql.db"),
	)

	var err error

	Engine, err = xorm.NewEngine("mysql", dns)

	if err != nil {
		zap.L().Error(fmt.Sprintf("数据库连接异常 %v", err))
	}

	Engine.ShowSQL(true)
	Engine.SetMapper(names.GonicMapper{})
}
