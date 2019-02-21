package phonenumber

import (
	"errors"
	"unicode"
)

// Number gets a phone number from a string
func Number(input string) (string, error) {

	finalStr := ""
	for _, char := range input {
		if unicode.IsDigit(char) {
			finalStr += string(char)
		}
	}

	if len(finalStr) == 11 {
		if string(finalStr[0]) != "1" {
			err := errors.New("First digit must be 1 in 11 digit number")
			return "", err
		}
		finalStr = finalStr[1:]
	}

	if len(finalStr) != 10 {
		err := errors.New("Too many or too few digits in phone number")
		return "", err
	}

	if string(finalStr[0]) == "0" || string(finalStr[0]) == "1" {
		err := errors.New("Area code cannot start with 0 or 1")
		return "", err
	}

	if string(finalStr[3]) == "0" || string(finalStr[3]) == "1" {
		err := errors.New("Exchange code cannot start with 0 or 1")
		return "", err
	}

	return finalStr, nil
}

// AreaCode gets the area code of a phone number
func AreaCode(input string) (string, error) {
	result, err := Number(input)
	if err != nil {
		return result, err
	}
	return result[:3], nil
}

// Format formats a phone number
func Format(input string) (string, error) {
	result, err := Number(input)
	if err != nil {
		return result, err
	}
	return "(" + result[:3] + ") " + result[3:6] + "-" + result[6:], nil
}
