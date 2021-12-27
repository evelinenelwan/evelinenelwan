package config

import (
	"api-gin/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := "root"
	password := "password"
	host := "tcp(127.0.0.1:3306)"
	database := "database_phone"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Phone{}, &models.ReviewRatingCategory{}, &models.User{})

	return db
}
