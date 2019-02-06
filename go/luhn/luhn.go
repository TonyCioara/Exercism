package luhn

import (
	"fmt"
	"unicode"
)

// Valid validates whether a credit card is valid
func Valid(input string) bool {

	if len(input) <= 1 {
		return false
	}

	readChar := false
	var allValues []int

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {

		char := runes[i]

		if !unicode.IsNumber(char) {
			continue
		}

		newValue := int(char) - 48
		fmt.Println("GOT HERE", newValue)
		if readChar == true {

			newValue *= 2
			if newValue > 9 {
				newValue -= 9
			}
		}

		allValues = append(allValues, newValue)

		readChar = !readChar
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
