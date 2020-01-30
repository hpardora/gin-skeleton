package database

import (
	"fmt"
	"gin-skeleton/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB

func init() {
	var err error
	address := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	DB, err = gorm.Open("mysql", user+":"+password+"@tcp("+address+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	// 	migrate
	DB.AutoMigrate(&models.User{})

	// 	显示详细日志
	DB.LogMode(true)
}
