package repository

import "codearena/internal/db/mysql"

func Migrate() {
	mysql.DBClient.AutoMigrate(&User{})
}
