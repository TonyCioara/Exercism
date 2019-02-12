package encode

import (
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) string {

	if len(input) == 0 {
		return ""
	}

	lastChar := input[0]
	lastCharCount := 1
	finalString := ""

	for i := 1; i < len(input); i++ {
		char := input[i]
		if char == lastChar {
			lastCharCount++
			continue
		}
		if lastCharCount == 1 {
			finalString += string(lastChar)
		} else {
			ct := strconv.Itoa(lastCharCount)
			finalString += ct + string(lastChar)
		}
		lastChar = char
		lastCharCount = 1
	}

	if lastCharCount == 1 {
		finalString += string(lastChar)
	} else {
		ct := strconv.Itoa(lastCharCount)
		finalString += ct + string(lastChar)
	}
	return finalString
}

func RunLengthDecode(input string) string {

	lastCount := ""
	finalString := ""
	for _, char := range input {

		if unicode.IsDigit(char) {
			lastCount += string(char)
			continue
		}

		str := string(char)

		if lastCount == "" {
			finalString += str
			continue
		}

		count, _ := strconv.Atoi(lastCount)

		for i := 0; i < int(count); i++ {
			finalString += str
		}
		lastCount = ""
	}
	return finalString
}
