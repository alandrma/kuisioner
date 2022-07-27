package database

import (
	"fmt"
	"kuisioner/config"
	"kuisioner/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "user_quiz"
const DB_PASSWORD = "KelapaIjo"
const DB_NAME = "quiz"
const DB_HOST = "10.9.20.30"
const DB_PORT = 3306

func InitDb() *gorm.DB {
	db := ConnectDB()
	db.AutoMigrate(&models.User{})
	return db
}

func ConnectDB() *gorm.DB {
	conf := config.DBConfig{
		Host:     DB_HOST,
		Port:     DB_PORT,
		User:     DB_USERNAME,
		DBName:   DB_NAME,
		Password: DB_PASSWORD,
	}

	db, err := gorm.Open(mysql.Open(conf.MysqlUrl()), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database: error=%v\n", err)
	}

	return db
}
