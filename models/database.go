package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "gorm_user:123@/snn_project?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Database connection failed!")
	}
	return db
}
