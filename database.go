package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// CreateConnection 创建数据库连接
func CreateConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, port, dbName,
		),
	)
}
