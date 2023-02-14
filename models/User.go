package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

// Creates new User
// takes username, password, returns ok bool
func (m Model) CreateUser(username string, password string) bool {
	db = m.Db
	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return false
	}

	newUser := User{Username: username, PasswordHash: string(passwordHashByte)}
	result := db.Create(&newUser)

	if result.Error != nil {
		return false
	}

	return true
}

// Checks if user exist
// returns ok bool, user User
func (m Model) GetUser(username, password string) (bool, User) {
	db = m.Db

	user := User{}

	result := db.First(&user, "username = ?", username)

	if result.Error != nil {
		return false, User{}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return false, User{}
	}

	return true, user
}
