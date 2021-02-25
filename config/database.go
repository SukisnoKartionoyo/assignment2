package config

import (
	"fmt"
	"log"
	"sesi7/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "go-test:test123@tcp(192.168.91.134:3306)/test-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(&models.Order{}, &models.Item{})
	fmt.Println("Connected to Database")

	DB = db
}
