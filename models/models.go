package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(32);not null"`
	Email    string `gorm:"type:varchar(64);not null;unique"`
	PassHash string `gorm:"type:char(32);not null"`
}

type UserRegister struct {
	Name           string `json:"name" form:"name" binding:"required"`
	Email          string `json:"email" form:"email" binding:"required"`
	Password       string `json:"password" form:"password" binding:"required"`
	RepeatPassword string `json:"rpassword" form:"rpassword" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
