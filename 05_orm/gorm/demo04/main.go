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

	// 按照主键更新整条记录
	//UpdateUserWhole(7)

	// 按主键只更新某些字段
	//UpdateUserFields(8)

	// 按照 Uid 更新某些字段
	UpdateByUID("901241512345600007", "李小明", "lixiaoming@gmail.com")
}

// UpdateUserWhole 按照主键更新整条记录
func UpdateUserWhole(id uint64) {
	user := User{
		Id:        id,
		Uid:       "901241512345600008",
		Name:      "赵飞飞",
		Email:     "zhaofeifei@qq.com",
		UpdatedAt: time.Now(),
	}

	result := db.Save(user)
	if result.Error != nil {
		fmt.Printf("整条更新失败: %v\n", result.Error)
	} else {
		fmt.Printf("整条更新成功，影响行数: %d\n", result.RowsAffected)
	}
}

// UpdateUserFields 按主键只更新某些字段
func UpdateUserFields(id uint64) {
	result := db.Model(&User{}).Where("id = ?", id).Updates(map[string]any{
		"name":  "李更新",
		"email": "liupdate@example.com",
	})
	// 也可以用结构体，但零值也会被更新（不推荐，除非你明确想设为空）
	// db.Model(&User{}).Where("id = ?", id).Updates(User{Name: "李更新", Email: "..."})

	if result.Error != nil {
		fmt.Printf("更新失败: %v\n", result.Error)
	} else if result.RowsAffected == 0 {
		fmt.Printf("要更新的用户 ID=%d 不存在\n", id)
	} else {
		fmt.Printf("字段更新成功！影响行数: %d\n", result.RowsAffected)
	}
}

// UpdateByUID 按 UID 更新
func UpdateByUID(uid, newName, newEmail string) {
	result := db.Model(&User{}).Where("uid = ?", uid).Updates(map[string]any{
		"name":  newName,
		"email": newEmail,
		// "updated_at": time.Now(), // GORM 会自动更新这个字段！不用写
	})

	if result.RowsAffected == 0 {
		fmt.Printf("UID=%s 不存在，更新失败\n", uid)
	} else {
		fmt.Printf("UID=%s 更新成功 → 姓名=%s, 邮箱=%s\n", uid, newName, newEmail)
	}
}
