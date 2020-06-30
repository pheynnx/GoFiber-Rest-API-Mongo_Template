package models

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
