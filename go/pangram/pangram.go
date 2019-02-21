package pangram

import (
	"unicode"
)

// IsPangram determines wether a string is a pangram or not
func IsPangram(input string) bool {

	letterMap := make(map[rune]bool)
	letterCount := 0

	for _, char := range input {

		if !unicode.IsLetter(char) {
			continue
		}

		char = unicode.ToLower(char)
		if letterMap[char] {
			continue
		}

		letterMap[char] = true
		letterCount++
	}

	if letterCount == 26 {
		return true
	}
	return false
}
