package models

import (
	"errors"

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

// type Url struct {
// 	gorm.Model
// 	ID           int    `json: "ID"`
// 	OriginalURL  string `json: "OriginalURL"`
// 	Token        string `json: "Token"`
// 	ShortenedURL string `json: "ShortenedURL"`
// }

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

// returns original url to redirect to
func (m Model) GetRedirect(token string) (bool, string) {
	db := m.Db

	var result Url

	db.Model(&Url{Token: token}).First(&result)

	OriginalURL := result.OriginalURL

	return true, OriginalURL

	// result = db.First(&Url{Token: token})

}
