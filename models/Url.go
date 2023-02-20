package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
	"krutki.pl/helpers"
)

type Url struct {
	gorm.Model
	ID          int
	OriginalURL string
	Token       string
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

// Checks if URL is already shortened
// returns (exists bool, token string)
func (m Model) UrlAlreadyExist(url string) (bool, string) {

	var result *Url
	var urlAlreadyExist bool

	db = m.Db

	_ = db.Model(&Url{}).Select("count(*) > 0").Where("original_url = ?", url).Find(&urlAlreadyExist).Error

	if urlAlreadyExist {

		err := db.Where("original_url = ?", url).First(&result).Error

		if err == nil {
			// return it only if url is present in the database
			return true, result.Token
		} else {
			// url is not present in db
			return false, ""
		}
	} else {
		// url is not present in database
		return false, ""
	}

}

// Creates new shortcut, new shortened URL
// returns (token string, error)
func (m Model) CreateNewShortcut(OriginalURL string) (string, error) {

	db := m.Db

	// check if url is already present

	urlAlreadyExist, existingToken := m.UrlAlreadyExist(OriginalURL)

	if urlAlreadyExist {
		return existingToken, nil
	}

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

func (m Model) GetAllLinks() []Url {
	var urls []Url

	db := m.Db

	result := db.Find(&urls)

	if result.Error != nil {
		panic("Url model: error querying db")
	}

	return urls
}
