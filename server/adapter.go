package server

import (
	"CodeArena/conf"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func init() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.V.GetString("mysql.user"),
		conf.V.GetString("mysql.pwd"),
		conf.V.GetString("mysql.host"),
		conf.V.GetInt("mysql.port"),
		conf.V.GetString("mysql.db"),
	)

	sqlx.NewMysql(dns)
}
