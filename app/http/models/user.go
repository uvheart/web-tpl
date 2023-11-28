package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username"`
	Email    string `gorm:"column:email"`
	Mobile   string `gorm:"column:mobile"`
}

func (u *User) TableName() string {
	return "user"
}
