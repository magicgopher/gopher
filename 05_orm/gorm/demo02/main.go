package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// 新增操作

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

	// 单条新增

	/*
		if err := InsertUser(); err != nil {
			fmt.Printf("单条插入失败: %v\n", err)
		} else {
			fmt.Println("单条插入成功！")
		}
	*/

	// 批量新增

	if err := InsertUsers(); err != nil {
		fmt.Printf("批量插入失败: %v\n", err)
	}

}

// InsertUser 单条新增
func InsertUser() error {
	user := User{
		Uid:   "901241512345600004",
		Name:  "王小明",
		Email: "wangxiaoming@example.com",
	}
	// Create 会自动填充 CreatedAt / UpdatedAt
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	// 打印自增ID
	fmt.Printf("新增成功，自增ID = %d, UID = %s\n", user.Id, user.Uid)
	return nil
}

// InsertUsers 批量新增
func InsertUsers() error {
	users := []User{
		{Uid: "901241512345600005", Name: "李小红", Email: "lixiaohong@example.com"},
		{Uid: "901241512345600006", Name: "赵铁柱", Email: "zhaotiezhu@example.com"},
		{Uid: "901241512345600007", Name: "钱多多", Email: "qianduoduo@example.com"},
	}
	result := db.CreateInBatches(users, 100) // 每100条一次提交，可自行调整
	if result.Error != nil {
		return result.Error
	}
	fmt.Printf("批量插入成功，共 %d 条\n", result.RowsAffected)
	return nil
}
