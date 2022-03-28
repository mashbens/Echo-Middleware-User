package config

import (
	"gorm-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:arif1412@tcp(127.0.0.1:3306)/alta_users?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())

	}

	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&models.User{})
}
