package main

import (
	"database/sql"
	"fmt"
	d3 "github.com/magicgopher/gopher/03_orm/01_gorm/demo04/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

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

	var err error

	// 打开数据库连接，返回 GORM 实例
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 获取底层 *sql.DB 对象（这是 Go 标准库的数据库连接池）
	// 用途：可以进行 Ping 测试、设置连接池参数、关闭连接等
	s, err := DB.DB()
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

	/*
		us := []d3.User{
			{Name: "张三", Age: 17, Sex: "男"},
			{Name: "李四", Age: 18, Sex: "女"},
			{Name: "萧火火", Age: 18, Sex: "男"},
			{Name: "林动", Age: 18, Sex: "男"},
			{Name: "萧薰儿", Age: 17, Sex: "女"},
		}

		err = CreateUsers(us)
		if err != nil {
			fmt.Printf("警告：批量插入部分失败: %v\n", err)
		}
		fmt.Println("批量插入成功！")
	*/

	/*
		// 模拟 100 条用户数据
		var users []d3.User
		for i := 1; i <= 100; i++ {
			users = append(users, d3.User{
				Name: fmt.Sprintf("User_%d", i),
				Age:  uint(rand.Intn(80) + 1),
				Sex: func(i int) string {
					if i%2 == 0 {
						return "男"
					} else {
						return "女"
					}
				}(i),
			})
		}

		// 调用分批插入函数，每批次插入 10 条
		err = CreateUsersInBatches(users, 10)

		// 在 main 中统一处理错误
		if err != nil {
			// 如果是核心业务数据插入失败，可以记录错误并退出
			log.Printf("程序在执行批量插入时遇到错误: %v", err)
			// os.Exit(1) // 如果需要强制退出可以取消注释
		}

		fmt.Println("所有分批任务执行完毕")
	*/

	user, err := FirstUser()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(user)

}

// CreateUsers 批量插入用户
func CreateUsers(users []d3.User) error {
	// 判断传入的内容
	if len(users) == 0 {
		return nil
	}

	//执行批量插入操作
	res := DB.Create(&users)

	// 错误处理
	if res.Error != nil {
		// 记录日志但不终止程序
		log.Printf("批量新增用户失败: %v", res.Error)
		return res.Error
	}

	// 打印受影响的行数
	fmt.Printf("成功插入 %d 条数据\n", res.RowsAffected)

	return nil
}

// CreateUsersInBatches 分批插入用户
func CreateUsersInBatches(users []d3.User, batchSize int) error {
	// 判断传入内容
	if len(users) == 0 {
		return nil
	}

	// result 是 *gorm.DB 类型，包含了 Error 和 RowsAffected
	res := DB.CreateInBatches(&users, batchSize)

	// 检查是否有错误
	if res.Error != nil {
		return fmt.Errorf("批量插入失败: %w", res.Error)
	}

	fmt.Printf("批量操作成功：共计划插入 %d 条，实际成功插入 %d 条\n", len(users), res.RowsAffected)
	return nil
}

// FirstUser 查找主键升序第一条记录
func FirstUser() (d3.User, error) {
	var user d3.User
	res := DB.First(&user)
	if res.Error != nil {
		log.Printf("查找主键升序第一条记录失败, 错误:%v\n", res.Error)
		return user, res.Error
	}
	return user, nil
}
