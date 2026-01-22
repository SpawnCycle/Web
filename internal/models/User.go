package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;type:varchar(20)"`
	Email    string `gorm:"unique;not null;type:varchar(50)"`
	Password string `gorm:"not null"`
}
