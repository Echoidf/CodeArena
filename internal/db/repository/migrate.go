package repository

import "content-assistant/internal/db/mysql"

func Migrate() {
	mysql.DBClient.AutoMigrate(&User{})
}
