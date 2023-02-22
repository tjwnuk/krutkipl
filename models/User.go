package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Ranks
// admin  - only admin, access to all sites
// mod 	  - moderator, access to selected sites
// reg    - regular user

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	PasswordHash string
	Rank         string
	gorm.Model
}

// Creates new User
// takes username, password, returns ok bool
func (m Model) CreateUser(username string, password string) bool {
	db = m.Db
	passwordHashByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return false
	}

	newUser := User{Username: username, PasswordHash: string(passwordHashByte), Rank: "reg"}
	result := db.Create(&newUser)

	if result.Error == nil {
		return true
	}

	return false
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

// Gets list of all users stored in database
// returns []Users slice
func (m Model) GetAllUsers() []User {
	var allUsers []User

	db := m.Db

	result := db.Find(&allUsers)

	if result.Error != nil {
		panic("User model: error querying db")
	}

	return allUsers
}

func (m Model) DeleteUser(idToDelete int) bool {

	db := m.Db

	db.Delete(&User{}, idToDelete)

	return true
}

// Grants mod privileges
// returns ok bool
func (m Model) GrantMod(userId int) bool {

	db := m.Db

	db.Model(&User{}).Where("id = ?", userId).Update("rank", "mod")

	return true
}

// Removes mod privileges
// returns ok bool
func (m Model) RemoveMod(userId int) bool {

	db := m.Db

	db.Model(&User{}).Where("id = ?", userId).Update("rank", "reg")

	return true
}
