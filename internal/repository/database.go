package repository

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

func NewGormDBConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_URL")))
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
