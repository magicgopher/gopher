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

	// 根据主键删除
	//DeleteByID(2)

	// 根据条件Uid删除
	//DeleteByUID("901241512345600004")

	// 根据多个Uid删除
	//BatchDeleteByUIDs([]string{"901241512345600001", "901241512345600003", "901241512345600005"})

	// 条件删除（比如删掉所有名字带“小”的）
	DeleteByCondition("%小%")
}

// DeleteByID 根据主键删除
func DeleteByID(id uint64) {
	result := db.Delete(&User{}, 2)
	if result.Error != nil {
		fmt.Printf("删除失败 Id=%d: %v\n", id, result.Error)
	}
	if result.RowsAffected == 0 {
		fmt.Printf("要删除的 ID=%d 不存在\n", id)
	} else {
		fmt.Printf("成功删除 ID=%d 的用户\n", id)
	}
}

// DeleteByUID 按业务 UID 删除
func DeleteByUID(uid string) {
	result := db.Where("uid = ?", uid).Delete(&User{})
	// RowsAffected = 0 → 你想操作的记录根本不存在
	// RowsAffected > 0 → 成功操作了这么多条
	if result.RowsAffected == 0 {
		fmt.Printf("UID=%s 不存在，无需删除\n", uid)
	} else {
		fmt.Printf("成功删除 UID=%s 的用户\n", uid)
	}
}

// BatchDeleteByUIDs 批量删除多个 UID
func BatchDeleteByUIDs(uids []string) {
	if len(uids) == 0 {
		return
	}
	result := db.Where("uid IN ?", uids).Delete(&User{})
	fmt.Printf("批量删除成功，共删除 %d 条（请求%d条，可能有不存在的）\n", result.RowsAffected, len(uids))
}

// DeleteByCondition 按条件删除（比如删名字带“小”的）
func DeleteByCondition(value string) {
	result := db.Where("name LIKE ?", value).Delete(&User{})
	if result.RowsAffected > 0 {
		fmt.Printf("条件删除成功：删除了 %d 个名字包含“小”的用户\n", result.RowsAffected)
	} else {
		fmt.Println("没有名字包含“小”的用户")
	}
}
