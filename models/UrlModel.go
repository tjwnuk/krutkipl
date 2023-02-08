package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"krutki.pl/helpers"
)

type Url struct {
	gorm.Model
	ID           int
	OriginalURL  string
	Token        string
	ShortenedURL string
}

// Checks if token exists in the database
// returns bool, true if exists or false if do not exist
func (m Model) TokenAlreadyExist(token string) bool {
	db = m.Db

	var tokenAlreadyExist bool

	// .Find() takes pointer, .Find(&variable)
	// otherwise do not work properly
	_ = db.Model(&Url{}).Select("count(*) > 0").Where("token = ?", token).Find(&tokenAlreadyExist).Error

	return tokenAlreadyExist
}

// Creates new shortcut, new shortened URL
// returns (token string, error)
func (m Model) CreateNewShortcut(OriginalURL string) (string, error) {

	db := m.Db

	// shortened url will have string like 19KKLP7O
	// tokenLength determines how many chars it will have
	var tokenLength int = 8
	token := helpers.GenerateToken(tokenLength)

	counter := 0
	for m.TokenAlreadyExist(token) {
		token = helpers.GenerateToken(tokenLength)

		counter++

		if counter > 30 {
			return "", errors.New("database is full")
		}
	}

	url := Url{OriginalURL: OriginalURL, Token: token}

	if r := db.Create(&url); r.Error != nil {
		return "", errors.New("error creating new entry")
	}

	return token, nil
}

// Get original url to redirect to
// returns ok bool, redirect url string
func (m Model) GetRedirectUrl(token string) (bool, string) {
	// get db object
	db := m.Db

	var result Url
	var OriginalURL string
	var resultString string

	// check if token exists in db
	// return OriginalURL if exist
	// or return false, "" if not
	err := db.Where("token = ?", token).First(&result).Error

	switch err {
	case nil:
		OriginalURL = result.OriginalURL
	case gorm.ErrRecordNotFound:
		return false, ""
	default:
		return false, ""
	}

	// add http prefixes to the URL to ensure it redirect correctly
	startsWithHttp := strings.HasPrefix(OriginalURL, "http://")
	startsWithHttps := strings.HasPrefix(OriginalURL, "https://")

	if !(startsWithHttp || startsWithHttps) {
		resultString = "http://" + OriginalURL
	} else {
		resultString = OriginalURL
	}

	return true, resultString

}
