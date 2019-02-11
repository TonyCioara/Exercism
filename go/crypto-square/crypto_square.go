package cryptosquare

import (
	"math"
	"unicode"
)

func Encode(input string) string {

	normalized := ""

	for _, char := range input {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			continue
		}
		normalized += string(unicode.ToLower(char))
	}

	rowLengthFloat := math.Sqrt(float64(len(normalized)))
	rowLength := int(rowLengthFloat)
	if rowLengthFloat > float64(rowLength) {
		rowLength++
	}

	var rows []string
	for i := 0; i < len(normalized); i += rowLength {
		j := i + rowLength
		if j > len(normalized) {
			j = len(normalized)
		}

		newRow := normalized[i:j]
		additions := rowLength - len(newRow)
		if additions > 0 {
			for i := 0; i <= additions; i++ {
				newRow += " "
			}
		}
		rows = append(rows, newRow)
	}

	finalString := ""
	for i := 0; i < rowLength; i++ {
		for j := 0; j < len(rows); j++ {
			finalString += string(rows[j][i])
		}
		if i < rowLength-1 {
			finalString += " "
		}
	}

	return finalString
}
