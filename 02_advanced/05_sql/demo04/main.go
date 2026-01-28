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
	if err := initDB(); err != nil {
		fmt.Printf("initDB() 失败:%v\n", err)
		return
	}
	// 一对多查询
	//oneToManyDemo()
	// 多对多查询
	manyToManyDemo()
}

// User 用户结构体
type User struct {
	ID     int64   // 主键
	Name   string  // 姓名
	Email  string  //邮件
	Orders []Order // 订单
}

// Order 订单结构体
type Order struct {
	ID      int     // 主键
	OrderNo string  // 订单编号
	Amount  float64 // 金额
}

// oneToManyDemo 多表查询（一对多）
func oneToManyDemo() {
	sqlStr := `SELECT u.id,
       u.name,
       u.email,
       o.id,
       o.order_no,
       o.amount
FROM users u
         LEFT JOIN
     orders o ON u.id = o.user_id
ORDER BY u.id`
	// 执行查询SQL语句
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("执行查询失败，SQL: %s, 错误: %v", sqlStr, err)
		return
	}
	// 释放资源
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("执行 rows.Close() 失败，SQL: %s, 错误: %v", sqlStr, err)
		}
	}()
	// 创建map来去重和聚合Users
	userMap := make(map[int64]*User)
	// 义切片保持顺序
	var userList []*User
	for rows.Next() {
		var uid, oid sql.NullInt64 // 使用 NullInt64 处理可能的 NULL 值（LEFT JOIN）
		var uName, uEmail, oNo sql.NullString
		var oAmount sql.NullFloat64
		// 扫描数据
		err := rows.Scan(&uid, &uName, &uEmail, &oid, &oNo, &oAmount)
		if err != nil {
			log.Printf("Scan 失败: %v\n", err)
			continue
		}
		// 1. 处理 User 部分
		userID := uid.Int64
		// 检查这个用户是否已经在 map 中存在
		currentUser, exists := userMap[userID]
		if !exists {
			// 如果不存在，创建新用户并加入 map 和 list
			currentUser = &User{
				ID:     uid.Int64,
				Name:   uName.String,
				Email:  uEmail.String,
				Orders: []Order{}, // 初始化切片
			}
			userMap[userID] = currentUser
			userList = append(userList, currentUser)
		}
		// 处理 Order 部分
		// 如果 oid 有效（即该行数据包含订单信息，不是 NULL）
		if oid.Valid {
			order := Order{
				ID:      int(oid.Int64),
				OrderNo: oNo.String,
				Amount:  oAmount.Float64,
			}
			// 将订单追加到当前用户的订单列表中
			currentUser.Orders = append(currentUser.Orders, order)
		}
	}
	// 打印结果
	for _, u := range userList {
		fmt.Printf("用户: %s (ID: %d)\n", u.Name, u.ID)
		if len(u.Orders) == 0 {
			fmt.Println("  无订单")
		}
		for _, o := range u.Orders {
			fmt.Printf("  - 订单号: %s, 金额: %.2f\n", o.OrderNo, o.Amount)
		}
	}
}

// Student 学生
type Student struct {
	ID      int64  // 对应 id
	SID     string // 对应 sid (学号)
	Name    string
	Courses []Course
}

// Course 课程
type Course struct {
	ID      int64  // 对应 id
	CID     string // 对应 cid (课程号)
	Title   string
	Credits int
}

// manyToManyDemo 多表查询（多对多）
func manyToManyDemo() {
	// 需要连接三张表：students -> student_courses -> courses
	sqlStr := `SELECT s.id,
       s.sid,
       s.name,
       c.id,
       c.cid,
       c.title,
       c.credits
FROM students s
         LEFT JOIN student_courses sc ON s.id = sc.student_id
         LEFT JOIN courses c ON sc.course_id = c.id
ORDER BY s.id`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("执行查询失败，SQL: %s, 错误: %v", sqlStr, err)
		return
	}
	// 释放资源
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("执行 rows.Close() 失败，SQL: %s, 错误: %v", sqlStr, err)
		}
	}()
	// 用于去重和聚合的 Map (Key: 学生的主键 ID)
	studentMap := make(map[int64]*Student)
	// 用于保持输出顺序的切片
	var studentList []*Student
	// 遍历结果集
	for rows.Next() {
		// 定义变量接收扫描结果
		// 学生字段 (Students 表是主表，理论上非空，但为了保险或通用性，除了ID外也可以用NullString)
		var sID int64
		var sSID, sName string
		// 课程字段 (Courses 表是被 LEFT JOIN 的，可能为 NULL，必须用 sql.NullXXX)
		var cID sql.NullInt64
		var cCID, cTitle sql.NullString
		var cCredits sql.NullInt32
		// Scan: 注意顺序要和 SQL SELECT 的顺序严格一致
		err := rows.Scan(
			&sID, &sSID, &sName, // 学生信息
			&cID, &cCID, &cTitle, &cCredits, // 课程信息 (可能为 NULL)
		)
		if err != nil {
			log.Printf("Scan 失败: %v\n", err)
			continue
		}
		// --- 逻辑 1: 处理学生 (聚合主项) ---
		// 检查该学生是否已存在于 Map 中
		currentStudent, exists := studentMap[sID]
		if !exists {
			// 如果是新遇到的学生，创建对象并初始化
			currentStudent = &Student{
				ID:      sID,
				SID:     sSID,
				Name:    sName,
				Courses: []Course{}, // 初始化空切片，避免 json 输出 null
			}
			// 存入 Map 和 List
			studentMap[sID] = currentStudent
			studentList = append(studentList, currentStudent)
		}
		// --- 逻辑 2: 处理课程 (聚合子项) ---
		// 只有当 cID 有效 (Valid = true) 时，才说明这一行包含了课程数据
		// 如果是"赵六"这种没课的学生，这里 cID.Valid 会是 false
		if cID.Valid {
			course := Course{
				ID:      cID.Int64,
				CID:     cCID.String,
				Title:   cTitle.String,
				Credits: int(cCredits.Int32),
			}
			// 将课程追加到当前学生的课程列表中
			currentStudent.Courses = append(currentStudent.Courses, course)
		}
	}
	// 检查遍历过程中是否有错
	if err = rows.Err(); err != nil {
		log.Printf("Rows 遍历错误: %v\n", err)
	}
	// --- 打印最终结果 ---
	for _, s := range studentList {
		fmt.Printf("学生: %s [%s] (ID:%d)\n", s.Name, s.SID, s.ID)
		if len(s.Courses) == 0 {
			fmt.Println("  (该学生暂未选修任何课程)")
		} else {
			for _, c := range s.Courses {
				fmt.Printf("  + 课程: %s [%s] - %d 学分\n", c.Title, c.CID, c.Credits)
			}
		}
		fmt.Println("- - - - -")
	}
}
