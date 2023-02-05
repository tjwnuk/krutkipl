package models

import (
	"errors"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// returns database object
func GetDB() (*gorm.DB, error) {

	err := godotenv.Load()

	if err != nil {
		panic("Error reading .env file")
	}

	dbfile := os.Getenv("SQLITE_FILENAME")

	if db == nil {
		db, err := gorm.Open(sqlite.Open(dbfile), &gorm.Config{})

		if err != nil {
			// panic("Database error")
			return nil, errors.New("an database error occured")
		}

		return db, nil
	}

	return db, nil
}

// returns project root
func GetProjectRootPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
