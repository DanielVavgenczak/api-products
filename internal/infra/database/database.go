package database

import (
	"fmt"
	"log"

	"github.com/DanielVavgenczak/api-products/internal/infra/entity"
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

func InitDB() *gorm.DB{
	db := connection()
	if err := migrations(db); err != nil {
		log.Fatal("mmigration error: ", err.Error())
	}
	return db
}

func connection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",DB_USERNAME,DB_PASSWORD,DB_HOST,DB_PORT,DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("database connection success...")
	return db
}

func migrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
	)
	if err != nil {
		return err
	}
	fmt.Println("migrations sucess...")
	return nil
}