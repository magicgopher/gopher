package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// 定义全局db变量
var db *sql.DB

func initDB() error {
	// dsn：Data Source Name
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
	// 设置与数据库建立连接的最大数目
	db.SetMaxOpenConns(100)
	// 最大空闲连接数
	db.SetMaxIdleConns(30)
	// 连接最大存活时间
	db.SetConnMaxLifetime(15 * time.Minute)
	// 连接连续空闲多久后被回收
	db.SetConnMaxIdleTime(90 * time.Second)
	return err
}

func closeDB() {
	if err := db.Close(); err != nil {
		log.Printf("db.Close 失败: %v\n", err)
	} else {
		log.Println("db.Close 成功！")
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB() 失败:%v\n", err)
		return
	}
	defer closeDB()
	// 查询一行数据
	//queryRowDemo()
	// 查询多行数据
	//queryMultiRowDemo()
	// 新增一行数据
	//insertRowDemo()
	// 更新一行数据
	//updateRowDemo()
	// 删除一行数据
	//deleteRowDemo()
	// 预编译查询
	//prepareQueryRows()
}

// User 用户结构体
type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

// queryRowDemo 查询单行数据
func queryRowDemo() {
	sqlStr := "SELECT * FROM user WHERE id=?"
	// 这个user结构体变量是用于封装查询结果的
	var u User
	err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age, &u.Sex)
	if err != nil {
		log.Printf("db.QueryRow() 失败: %v\n", err)
		return
	}
	fmt.Printf("user{id:%d name:%s age:%d sex: %s}\n", u.Id, u.Name, u.Age, u.Sex)
}

// queryMultiRowDemo 查询多行数据
func queryMultiRowDemo() {
	sqlStr := "SELECT * FROM user WHERE id>?"
	rows, err := db.Query(sqlStr, 10)
	if err != nil {
		log.Printf("db.Query() 失败: %v\n", err)
		return
	}
	// rows关闭
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("rows.Close() 失败: %v\n", err)
		} else {
			log.Println("rows.Close() 成功.")
		}
	}()
	// 遍历查询的rows
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Sex); err != nil {
			log.Printf("rows.Scan() 失败: %v\n", err)
			return
		}
		fmt.Printf("user{id:%d name:%s age:%d sex: %s}\n", u.Id, u.Name, u.Age, u.Sex)
	}
}

// insertRowDemo 新增一行数据
func insertRowDemo() {
	sqlStr := "INSERT INTO `user` (`name`, `age`, `sex`) VALUE (?,?,?)"
	ret, err := db.Exec(sqlStr, "萧火火", 18, "男")
	if err != nil {
		log.Printf("db.Exec() INSERT ROW 失败: %v\n", err)
		return
	}
	insertId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("获取新插入的数据Id失败:%v\n", err)
		return
	}
	fmt.Printf("插入成功, 新插入数据的id: %d.\n", insertId)
}

// updateRowDemo 更新一行数据
func updateRowDemo() {
	sqlStr := "UPDATE user SET name=?, age=?, sex=? WHERE id=?"
	ret, err := db.Exec(sqlStr, "萧鑫", 20, "女", 20)
	if err != nil {
		log.Printf("db.Exec() UPDATE ROW 失败: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("获取更新操作影响的行数失败: %v\n", err)
		return
	}
	fmt.Printf("更新操作影响的行数: %d\n", n)
}

// deleteRowDemo 删除一行数据
func deleteRowDemo() {
	sqlStr := "DELETE FROM user WHERE id=?"
	ret, err := db.Exec(sqlStr, 21)
	if err != nil {
		log.Printf("db.Exec() DELETE ROW 失败: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("获取删除操作影响的行数失败: %v\n", err)
		return
	}
	fmt.Printf("删除操作影响的行数: %d\n", n)
}

func prepareQueryRows() {
	sqlStr := "SELECT * FROM user WHERE age=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("db.Prepare() 失败: %v\n", err)
		return
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			log.Printf("stmt.Close() 关闭失败: %v\n", closeErr)
		} else {
			log.Println("stmt.Close() 关闭成功.")
		}
	}()
	rows, err := stmt.Query(20)
	if err != nil {
		log.Printf("db.Query() 失败: %v\n", err)
		return
	}
	// rows关闭
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("rows.Close() 失败: %v\n", err)
		} else {
			log.Println("rows.Close() 成功.")
		}
	}()
	// 遍历查询的rows
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Sex); err != nil {
			log.Printf("rows.Scan() 失败: %v\n", err)
			return
		}
		fmt.Printf("user{id:%d name:%s age:%d sex: %s}\n", u.Id, u.Name, u.Age, u.Sex)
	}
}
