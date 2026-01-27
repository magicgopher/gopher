package main

import (
	"database/sql"
	"errors"
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
	// 事务操作
	//transactionDemo()
	// 事务操作使用预编译
	transactionPrepareDemo()
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

// prepareQueryRows 预编译查询
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

// transactionDemo 事务操作示例
func transactionDemo() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Printf("开启事务失败: %v\n", err)
		return
	}
	// 延迟回滚机制
	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Printf("事务回滚失败: %v\n", err)
		}
	}()
	// 模拟转账逻辑的变量
	user1 := 1
	user2 := 2
	money := 500.00
	// 用户1执行扣款操作
	fmt.Printf("从用户 %d 扣减 %.2f 元\n", user1, money)
	res1, err := tx.Exec("UPDATE account SET balance = balance - ? WHERE user_id = ?", money, user1)
	if err != nil {
		log.Printf("扣款 SQL 执行错误: %v\n", err)
		return // 触发 defer 回滚
	}
	// 校验：是否真的扣到了钱
	rows1, err := res1.RowsAffected()
	if err != nil {
		log.Printf("获取扣款影响行数失败: %v\n", err)
		return
	}
	if rows1 != 1 {
		log.Printf("扣款失败: 找不到用户 %d 或其他原因\n", user1)
		return // 触发 defer 回滚
	}
	//fmt.Println("系统即将发生严重崩溃 (Panic) ...")
	//panic("模拟突发崩溃：机房停电了 / 代码原本有 Bug / 内存溢出")
	// 用户2入账操作
	fmt.Printf("给用户 %d 增加 %.2f 元\n", user2, money)
	res2, err := tx.Exec("UPDATE account SET balance = balance + ? WHERE user_id = ?", money, user2)
	if err != nil {
		log.Printf("入账 SQL 执行错误: %v\n", err)
		return // 触发 defer 回滚
	}
	// 处理 RowsAffected 的 error
	rows2, err := res2.RowsAffected()
	if err != nil {
		log.Printf("获取入账影响行数失败: %v\n", err)
		return // 触发 defer 回滚
	}
	// 业务逻辑校验
	if rows2 != 1 {
		log.Printf("入账失败: 找不到用户 %d\n", user2)
		return // 触发 defer 回滚
	}
	// 提交事务
	// 只有执行到这里，数据库才会真正变更
	if err := tx.Commit(); err != nil {
		log.Printf("事务提交失败: %v\n", err)
		return
	}
	fmt.Println("转账成功！")
}

// transactionPrepareDemo 事务操作示例（使用预编译）
func transactionPrepareDemo() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Printf("开启事务失败: %v\n", err)
		return
	}
	// 延迟回滚
	defer func() {
		err := tx.Rollback()
		// 如果错误是 sql.ErrTxDone，说明事务已经 Commit 成功了，不需要报错
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			log.Printf("严重警告: 事务回滚失败: %v\n", err)
		}
	}()
	// 模拟转账数据
	user1 := 1
	user2 := 2
	money := 500.00
	// 转账 SQL 使用加法，扣款传负数，入账传正数
	sqlStr := "UPDATE account SET balance = balance + ? WHERE user_id = ?"
	// 使用 tx.Prepare 预编译SQL
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		log.Printf("预编译 SQL 失败: %v\n", err)
		return // 触发 defer tx.Rollback
	}
	// 延迟关闭 Stmt
	// Stmt 也是数据库资源，必须关闭
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("警告: 预编译语句关闭失败: %v\n", err)
		}
	}()
	// user1扣款操作
	fmt.Printf("Step 1: 从用户 %d 扣减 %.2f 元\n", user1, money)
	// 执行：传入负数 (-money)
	res1, err := stmt.Exec(-money, user1)
	if err != nil {
		log.Printf("扣款 SQL 执行错误: %v\n", err)
		return // 触发 defer
	}
	// 检查 RowsAffected 及其错误
	rows1, err := res1.RowsAffected()
	if err != nil {
		log.Printf("获取扣款影响行数时出错: %v\n", err)
		return // 触发 defer
	}
	if rows1 != 1 {
		log.Printf("扣款业务失败: 未找到用户 %d 或未发生更新\n", user1)
		return // 触发 defer
	}
	fmt.Println("系统即将发生严重崩溃 (Panic) ...")
	panic("模拟突发崩溃：机房停电了 / 代码原本有 Bug / 内存溢出")
	// user2入账操作
	fmt.Printf("Step 2: 给用户 %d 增加 %.2f 元\n", user2, money)
	// 传入正数 (money) - 复用同一个 stmt
	res2, err := stmt.Exec(money, user2)
	if err != nil {
		log.Printf("入账 SQL 执行错误: %v\n", err)
		return // 触发 defer
	}
	// 检查 RowsAffected 及其错误
	rows2, err := res2.RowsAffected()
	if err != nil {
		log.Printf("获取入账影响行数时出错: %v\n", err)
		return // 触发 defer
	}
	if rows2 != 1 {
		log.Printf("入账业务失败: 未找到用户 %d\n", user2)
		return // 触发 defer
	}
	// 提交事务
	if err := tx.Commit(); err != nil {
		log.Printf("事务提交失败: %v\n", err)
		return
	}
	fmt.Println("转账成功")
}
