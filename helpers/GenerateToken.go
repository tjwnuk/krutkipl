package helpers

import (
	"math/rand"
	"time"
)

// Generates token for shortened URLs
// it will appear in shortened URL, like krutki.pl/l/AABBCC
// takes: (int) length of the string
// returns: (string) generated token
func GenerateToken(length int) string {
	var token_byte_arr []rune
	var token_string string

	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	rand.Seed(time.Now().Unix())

	//generate N runes for specified token length
	for i := 0; i < length; i++ {
		var randomChar rune

		// random choose from digit and letter
		if rand.Intn(2) == 1 {
			// choose random capital letter
			randomChar = 'A' + rune(rand.Intn(26))
		} else {
			// choose random digit
			randomChar = digits[rand.Intn(10)]
		}
		token_byte_arr = append(token_byte_arr, randomChar)
	}

	token_string = string(token_byte_arr)
	return token_string
}
