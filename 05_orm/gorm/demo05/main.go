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
	//QueryAll()

	// 根据主键查询
	//QueryByID(7)

	// 根据Uid查询
	//QueryByUID("901241512345600007")

	// 模糊搜索（名字包含某个字）
	//QueryLikeName("飞")

	// 多条件组合查询（AND）
	//QueryMulti()

	// 分页加排序
	//QueryPage(1, 10) // 第一页，每页10条数据
	//QueryPage(2, 10) // 第一页，每页10条数据
	//QueryPage(3, 10) // 第一页，每页10条数据

	// 只查部分字段
	//QuerySelect()

	// 查询总记录数
	//CountTotal()

	//
	ExistsByUID("901241512345600031")
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

// QueryByUID 根据Uid查询
func QueryByUID(uid string) *User {
	var user User
	if err := db.Where("uid = ?", uid).First(&user).Error; err != nil {
		fmt.Printf("UID=%s 不存在\n", uid)
		return nil
	}
	fmt.Printf("找到用户 → %s (%s)\n", user.Name, user.Email)
	return &user
}

// QueryLikeName 根据名字模糊查询
func QueryLikeName(namePart string) {
	var users []User
	db.Where("name LIKE ?", "%"+namePart+"%").Find(&users)
	fmt.Printf("名字包含 [%s] 的用户有 %d 个：\n", namePart, len(users))
	for _, u := range users {
		fmt.Printf("  → %s (UID=%s)\n", u.Name, u.Uid)
	}
}

// QueryMulti 多条件组合查询 AND
func QueryMulti() {
	var users []User
	db.Where("name LIKE ?", "张%").
		Where("email LIKE ?", "%@example.com").
		Find(&users)
	fmt.Printf("同时满足“名字带张 + 邮箱含example”的有 %d 条\n", len(users))

	fmt.Println("信息如下：")

	for _, u := range users {
		fmt.Printf("Uid:%s, 姓名:%s, 邮件:%s\n", u.Uid, u.Name, u.Email)
	}
}

// QueryPage 分页加排序
func QueryPage(page, pageSize int) {
	var users []User
	offset := (page - 1) * pageSize
	db.Order("id DESC").
		Offset(offset).
		Limit(pageSize).Find(&users)

	fmt.Printf("第 %d 页（每页 %d 条）共 %d 条：\n", page, pageSize, len(users))
	for _, u := range users {
		fmt.Printf("  [%d] %s - %s - %s\n", u.Id, u.Uid, u.Name, u.Email)
	}
}

// QuerySelect 只查部分字段（提升性能 + 减少内存）
func QuerySelect() {
	var results []struct {
		Id    uint64
		Name  string
		Email string
	}
	db.Model(&User{}).Select("id", "name", "email").Limit(5).Find(&results)
	fmt.Println("只查 id、name、email 字段的前5条：")
	for _, r := range results {
		fmt.Printf("  %d | %s | %s\n", r.Id, r.Name, r.Email)
	}
}

// CountTotal 统计总数
func CountTotal() {
	var count int64
	db.Model(&User{}).Count(&count)
	fmt.Printf("用户表当前总人数：%d\n", count)
}

// ExistsByUID 判断是否存在
func ExistsByUID(uid string) bool {
	var count int64
	db.Model(&User{}).Where("uid = ?", uid).Count(&count)
	exists := count > 0
	fmt.Printf("UID=%s %s\n", uid, map[bool]string{true: "已存在", false: "不存在"}[exists])
	return exists
}
