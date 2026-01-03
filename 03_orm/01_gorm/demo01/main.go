package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 使用原生 database/sql 操作mysql数据库

// User 用户结构体对应 tb_user 表
type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

// validateUser 辅助函数，验证User结构体的内容
func validateUser(u User) error {
	if u.Name == "" {
		return fmt.Errorf("用户名称不能为空")
	}

	if u.Age <= 0 {
		return fmt.Errorf("年龄必须大于0")
	}

	if u.Gender < 0 || u.Gender > 2 {
		return fmt.Errorf("性别值无效，必须为 0(未知)、1(男) 或 2(女)")
	}

	return nil
}

func insertUser(db *sql.DB, u User) (User, error) {
	// 验证
	if err := validateUser(u); err != nil {
		log.Printf("用户数据校验失败: %v", err)
		return User{}, err
	}

	// 执行SQL
	sql := "INSERT INTO tb_user (name, age, gender) VALUES (?, ?, ?)"
	res, err := db.Exec(sql,
		u.Name,
		u.Age,
		u.Gender,
	)

	// 获取自增ID
	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("获取自增ID失败: %v", err)
		return User{}, err
	}

	// 构造并返回完整对象
	newUser := User{
		ID:     id,
		Name:   u.Name,
		Age:    u.Age,
		Gender: u.Gender,
	}

	// 返回新增的用户信息
	return newUser, nil
}

// updateUser 更新用户操作
func updateUser(db *sql.DB, u User) (User, error) {
	if u.ID <= 0 {
		return User{}, fmt.Errorf("用户ID无效")
	}

	// 业务规则校验
	if err := validateUser(u); err != nil { // false 表示 name 可选
		log.Printf("更新用户数据校验失败: %v", err)
		return User{}, err
	}

	// 执行更新（只更新 name, age, gender）
	res, err := db.Exec(`
        UPDATE tb_user 
        SET name = ?, age = ?, gender = ? 
        WHERE id = ?`,
		u.Name, u.Age, u.Gender, u.ID,
	)
	if err != nil {
		log.Printf("更新用户数据库错误: %v", err)
		return User{}, err
	}

	// 检查是否真的更新了一行
	affected, err := res.RowsAffected()
	if err != nil {
		log.Printf("获取影响行数失败: %v", err)
		return User{}, err
	}
	if affected == 0 {
		return User{}, fmt.Errorf("未找到 ID 为 %d 的用户，更新失败", u.ID)
	}

	// 5. 返回更新后的完整对象（直接使用传入的 u，因为字段已更新）
	return u, nil
}

// 查询全部用户
func listAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, age, gender FROM tb_user ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("关闭数据库结果集失败: %v", closeErr)
		}
	}()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Gender); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func main() {
	// 数据源
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		"root",
		"12345678",
		"tcp",
		"127.0.0.1",
		"13306",
		"test_db",
	)

	// 这里并不会去建立和数据库的实际的TCP连接
	// 返回一个 *sql.DB实例
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("dsn数据源格式错误: %v", err)
	}

	//
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据连接失败: %v\n", err)
	}
	fmt.Println(db)
	fmt.Println("数据连接成功！")

	// defer关闭数据库连接
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("数据库连接关闭失败: %v\n", err)
		} else {
			log.Println("数据连接关闭成功！")
		}
	}(db)

	// 新增操作
	/*
		user1 := User{
			Name:   "萧炎",
			Age:    18,
			Gender: 1,
		}
		created, err := insertUser(db, user1)
		if err != nil {
			log.Printf("插入失败: %v", err)
		} else {
			fmt.Printf("插入成功: %+v\n", created)
		}

		badUser := User{
			Name:   "", // 空名称
			Age:    -5, // 负年龄
			Gender: 5,  // 无效性别
		}
		_, err = insertUser(db, badUser)
		if err != nil {
			fmt.Printf("预期校验错误: %v\n", err)
		}
	*/

	// 更新示例
	/*
		updateInput := User{
			ID:     1,
			Name:   "赵六", // 可以改名
			Age:    35,
			Gender: 1,
		}
		updatedUser, err := updateUser(db, updateInput)
		if err != nil {
			log.Printf("更新失败: %v", err)
		} else {
			fmt.Printf("更新成功: %+v\n", updatedUser)
		}

		// 错误示例：年龄非法
		badUpdate := User{
			ID:     1,
			Name:   "测试",
			Age:    0, // 非法
			Gender: 1,
		}
		_, err = updateUser(db, badUpdate)
		if err != nil {
			fmt.Printf("预期校验错误: %v\n", err) // 输出：年龄必须大于0
		}
	*/

	// 查询所有用户
	/*
		users, err := listAllUsers(db)
		if err != nil {
			log.Printf("查询所有用户失败: %v\n", err)
		}
		for _, u := range users {
			fmt.Println(u)
		}
	*/
}
