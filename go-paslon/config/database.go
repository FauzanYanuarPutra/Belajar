package config

import (
	"go-paslon/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
		// dsn := "root:@tcp(localhost:3306)/go_paslon?charset=utf8mb4&parseTime=True&loc=Local"
    // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		dsn := "user=postgres password=root dbname=go_paslon sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

		db.AutoMigrate(&models.Paslon{})
		
    DB = db
}


