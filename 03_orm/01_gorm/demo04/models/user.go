package models

var TableName = "tb_user"

type User struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
	Age  uint   `json:"age" gorm:"column:age"`
	Sex  string `json:"sex" gorm:"column:sex"`
}

func (User) TableName() string {
	return TableName
}
