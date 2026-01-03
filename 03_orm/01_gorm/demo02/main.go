package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	// 数据源
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s&timeout=%s",
		"root",      //用户名
		"12345678",  //密码
		"tcp",       //协议
		"127.0.0.1", //IP地址或者域名
		"13306",     //端口号
		"test_db",   //使用的数据库
		"utf8mb4",   //字符集
		"True",      //是否解析时间类型
		"Local",     //时区
		"10s",       //连接超时
	)

	// 打开数据库连接，返回 GORM 实例
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 获取底层 *sql.DB 对象（这是 Go 标准库的数据库连接池）
	// 用途：可以进行 Ping 测试、设置连接池参数、关闭连接等
	s, err := db.DB()
	if err != nil {
		log.Fatal("获取底层数据库连接池失败:", err)
	}

	// 发送 Ping 请求，测试数据库连接是否正常
	// 如果数据库不可达、网络中断或连接已断开，这里会返回错误
	err = s.Ping()
	if err != nil {
		log.Fatal("数据库 Ping 失败，连接不可用:", err)
	}

	//使用 defer 延迟执行关闭操作，确保函数退出前释放数据库连接资源
	defer func(s *sql.DB) {
		//关闭数据库连接，并释放所有正在使用的资源
		err := s.Close()
		if err != nil {
			//数据库连接关闭失败，记录错误日志
			log.Printf("关闭数据库连接时发生错误: %v", err)
		} else {
			//成功关闭，打印信息
			fmt.Println("数据库连接已成功关闭")
		}
	}(s)

	// 数据库连接成功后输出提示
	fmt.Println("数据库连接成功！")

	// 设置空闲连接池中的最大连接数（建议 10 左右）
	s.SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数（包括使用中 + 空闲）
	s.SetMaxOpenConns(100)
	// 设置连接最大存活时间（防止连接长期存在导致问题，建议 1~4 小时）
	s.SetConnMaxLifetime(time.Hour * 4)
}
