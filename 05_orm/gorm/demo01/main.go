package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB // 全局 db，方便后面调用

func main() {
	// dsn：数据源
	dsn := "host=localhost user=root password=12345678 dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	// 连接postgresql数据库
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	fmt.Println("数据库连接成功！")
}
