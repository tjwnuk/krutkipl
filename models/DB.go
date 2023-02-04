package models

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"os"
)

var db *gorm.DB

// returns database object
func GetDB() (*gorm.DB, error) {
	if db == nil {
		db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

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
