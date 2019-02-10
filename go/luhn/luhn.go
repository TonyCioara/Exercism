package luhn

import (
	"unicode"
)

// Valid validates whether a credit card is valid
func Valid(input string) bool {

	readChar := false
	var allValues []int

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {

		char := runes[i]

		if string(char) == " " {
			continue
		}

		if !unicode.IsNumber(char) {
			return false
		}

		newValue := int(char) - 48
		if readChar == true {
			newValue *= 2
			if newValue > 9 {
				newValue -= 9
			}
		}

		allValues = append(allValues, newValue)

		readChar = !readChar
	}

	if len(allValues) == 1 {
		return false
	}

	sum := 0
	for _, value := range allValues {
		sum += int(value)
	}

	if sum%10 == 0 {
		return true
	}
	return false

}
