package prime

import "math"

// Factors finds all the prime factors of a number
func Factors(num int64) []int64 {

	returnArray := []int64{}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		i64 := int64(i)
		if num%i64 == 0 {
			returnArray = append(returnArray, i64)
			num /= i64
			i = 1
		}
	}
	if num != 1 {
		returnArray = append(returnArray, num)
	}
	return returnArray
}
