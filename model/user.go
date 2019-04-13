package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	UserId   int    `gorm:"user_id" json:"user_id"`
	UserName string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}
