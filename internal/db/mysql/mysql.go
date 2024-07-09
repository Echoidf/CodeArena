package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func init() {
	var err error
	// dsn := "root:admin123@tcp(127.0.0.1:3306)/codearena?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:admin123@unix(/tmp/mysql.sock)/codearena?charset=utf8mb4&parseTime=True&loc=Local"
	DBClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
