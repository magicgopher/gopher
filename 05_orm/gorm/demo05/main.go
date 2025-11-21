package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB // 定义全局DB

// User 用户模型
type User struct {
	Id        uint64    `gorm:"primaryKey;column:id"`
	Uid       string    `gorm:"type:varchar(64);uniqueIndex;column:uid"`
	Name      string    `gorm:"type:varchar(50);column:name"`
	Email     string    `gorm:"type:varchar(100);column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName 表名默认就是 tb_user
func (User) TableName() string {
	return "tb_user"
}

func main() {
	// 数据源
	dsn := "host=127.0.0.1 user=root password=12345678 dbname=test_db port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	fmt.Println("数据库连接成功！")

	// 查询操作
	QueryAll()
}

// QueryAll 查所有（最简单）
func QueryAll() {
	var users []User
	db.Find(&users) // 注意：Find 能接收 slice 和单个结构体
	fmt.Printf("共 %d 条记录：\n", len(users))
	for _, u := range users {
		fmt.Printf("  ID=%d | UID=%s | 姓名=%s | 邮箱=%s\n", u.Id, u.Uid, u.Name, u.Email)
	}
}

// QueryByID 根据主键查询
func QueryByID(id uint64) {
	var user User
	err := db.First(&user, id).Error // First 按主键顺序查第一个
	if err != nil {
		fmt.Printf("ID=%d 不存在或查询出错\n", id)
		return
	}
	fmt.Printf("查到用户 → ID=%d, 姓名=%s, 邮箱=%s\n", user.Id, user.Name, user.Email)
}
