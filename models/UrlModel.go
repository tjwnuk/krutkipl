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

// Creates new shortcut, new shortened URL
// returns (ok bool, ShortenedUrl string)
func (m Model) CreateNewShortcut(OriginalURL string) (string, error) {

	db := m.Db

	token := helpers.GenerateToken(8)
	shortenedUrl := "krutki.pl/l/" + token

	url := Url{OriginalURL: OriginalURL, Token: token, ShortenedURL: shortenedUrl}

	if result := db.Create(&url); result.Error != nil {
		return "", errors.New("error creating new entry")
	}

	return shortenedUrl, nil
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
