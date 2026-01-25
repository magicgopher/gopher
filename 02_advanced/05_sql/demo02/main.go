package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 使用原生的database/sql标准库操作数据库
// 简单增删改查操作

var db *sql.DB

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:13306)/test_db"
	var err error
	// Open打开 创建连接池管理器 + DSN校验格式
	db, err = sql.Open("mysql", dsn)
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

	// 查询全部
	//findUsers()

	// 查询单行
	//findByUserOne()

	// 新增操作
	//createUser("小哈哈", 16, 1)

	// 更新操作
	//updateUser(5, "大哈哈", 19, 2)

	// 删除操作
	//deleteUser(5)
}

type User struct {
	ID     int
	Name   string
	Age    int
	Gender int
}

// findUsers 查询多行
func findUsers() {
	// 执行 SELECT 查询
	rows, err := db.Query("SELECT * FROM tb_user")
	if err != nil {
		log.Printf("db.Query() 失败: %v\n", err)
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("rows.Close() 关闭失败: %v\n", err)
		}
	}()
	// 用来存放查询到的所有用户记录
	var users []User
	// 逐行读取查询结果
	for rows.Next() {
		// 每次循环创建一个新的 User 结构体变量，用于接收当前这一行的数据
		var user User
		// 将当前行的字段值扫描（映射）到 user 结构体的对应字段
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender); err != nil {
			log.Printf("rows.Scan() 失败: %v\n", err)
		}
		// 把扫描成功的用户添加到切片中
		users = append(users, user)
	}
	// 遍历所有查询到的用户，并打印到控制台
	for _, user := range users {
		fmt.Println(user)
	}
}

// findByOne 查询单个用户
func findByUserOne() {
	// 定义用于存储查询到的结果
	var user User
	// 查询单行记录
	// 标准库 database/sql 的 Scan 不支持直接传结构体
	if err := db.QueryRow("SELECT * FROM tb_user WHERE id=?", 1).Scan(&user.ID, &user.Name, &user.Age, &user.Gender); err != nil {
		log.Printf("db.QueryRow() 失败: %v\n", err)
		return
	}
	fmt.Println(user)
}

// createUser 新增操作
func createUser(name string, age, gender int) {
	result, err := db.Exec("INSERT INTO tb_user (name, age, gender) VALUES (?, ?, ?)", name, age, gender)
	if err != nil {
		log.Printf("db.Exec() 新增操作失败: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败: %v\n", err)
		return
	}
	fmt.Printf("新增用户成功，受影响行数: %d\n", affected)
}

// updateUser 更新一条用户记录
func updateUser(id int64, newName string, newAge int, newGender int) {
	// 准备要更新的字段
	// 这里假设 updated_at 字段会自动更新（如果表有 ON UPDATE CURRENT_TIMESTAMP）
	result, err := db.Exec(`UPDATE tb_user SET name = ?, age = ?, gender = ? WHERE id = ?`, newName, newAge, newGender, id)
	if err != nil {
		log.Printf("更新用户失败 id=%d : %v", id, err)
		return
	}
	// 获取受影响的行数（正常情况下应该 == 1）
	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败 id=%d : %v (但更新大概率已成功)", id, err)
		fmt.Printf("用户 id=%d 已尝试更新，但无法确认影响行数\n", id)
		return
	}
	if affected == 0 {
		log.Printf("更新操作执行了，但实际未影响任何行 id=%d （可能记录不存在）", id)
		fmt.Printf("未找到 id=%d 的用户，更新未生效\n", id)
		return
	}
	if affected > 1 {
		log.Printf("警告：更新影响了多行 id=%d ，实际影响 %d 行（可能有问题）", id, affected)
	}
	fmt.Printf("更新用户成功 id=%d ，受影响行数: %d\n", id, affected)
}

// deleteUser 删除一条用户记录
func deleteUser(id int64) {
	result, err := db.Exec("DELETE FROM tb_user WHERE id = ?", id)
	if err != nil {
		log.Printf("删除用户失败 id=%d : %v", id, err)
		return
	}
	// 获取受影响的行数（正常情况下应该 == 1）
	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败 id=%d : %v (但删除大概率已成功)", id, err)
		fmt.Printf("用户 id=%d 已尝试删除，但无法确认影响行数\n", id)
		return
	}
	if affected == 0 {
		log.Printf("删除操作执行了，但实际未影响任何行 id=%d （可能记录本来就不存在）", id)
		fmt.Printf("未找到 id=%d 的用户，无需删除\n", id)
		return
	}
	if affected > 1 {
		log.Printf("警告：删除影响了多行 id=%d ，实际影响 %d 行（可能有问题）", id, affected)
	}
	fmt.Printf("删除用户成功 id=%d ，受影响行数: %d\n", id, affected)
}
