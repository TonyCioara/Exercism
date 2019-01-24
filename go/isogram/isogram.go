package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			if letters[unicode.ToLower(letter)] {
				return false
			}
			letters[unicode.ToLower(letter)] = true
		}
	}
	return true
}