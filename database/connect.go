package database

import (
	"ericarthurc/fiberAPI/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to database
func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", os.Getenv("DATABASE"))
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated")
}
