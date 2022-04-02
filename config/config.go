package config

import (
	"fmt"
	"gorm-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_USER = "root"
const DB_PASS = "arif1412"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "test_user"

const DB_USER_TEST = "root"
const DB_PASS_TEST = "arif1412"
const DB_HOST_TEST = "127.0.0.1"
const DB_PORT_TEST = "3306"
const DB_NAME_TEST = "test_user"

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
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

func DBManager() *gorm.DB {
	return DB
}

func InitDBTest() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
}
