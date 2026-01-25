package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 使用原生的database/sql标准库操作数据库
// 连接数据库/关闭数据库连接操作

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:13306)/test_db"
	// Open打开 创建连接池管理器 + DSN校验格式
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("sql.Open 失败（DSN 格式/驱动问题）: %v", err)
	}
	// Ping() 现在才真正尝试建立连接、登录、验证权限
	if err = db.Ping(); err != nil {
		log.Fatalf("数据库连接失败（账号、密码、地址、端口、数据库名任一错误）: %v", err)
	}
	fmt.Println("数据库连接成功！")
	// defer关闭连接
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("db.Close 失败: %v\n", err)
		} else {
			log.Println("db.Close 成功！")
		}
	}()
}
