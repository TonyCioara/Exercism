package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(num int) bool {

	runes := []rune(strconv.Itoa(num))

	power := float64(len(runes))

	var productSum float64
	for _, char := range runes {
		str := string(char)
		digit, _ := strconv.Atoi(str)

		productSum += math.Pow(float64(digit), power)
	}
	return int(productSum) == num
}
