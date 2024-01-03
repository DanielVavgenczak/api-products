package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "docker"
	DB_NAME = "db_api_product"
	DB_PORT = "3306"
	DB_HOST = "127.0.0.1"
)

var db *gorm.DB

func InitDB() *gorm.DB{
	db := Connection()
	return db
}

func Connection() (*gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",DB_USERNAME,DB_PASSWORD,DB_HOST,DB_PORT,DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("database connection success...")
	return db
}