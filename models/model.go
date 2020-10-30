package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/open_9lan_cn?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
