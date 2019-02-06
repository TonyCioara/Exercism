package summultiples

// SumMultiples calculates the sum of all possible multiples of the divisors, smaller than the limit
func SumMultiples(limit int, divisors []int) int {
	resultMap := make(map[int]bool)
	sum := 0

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for currentNum := divisor; currentNum < limit; currentNum += divisor {
			if resultMap[currentNum] {
				continue
			}
			resultMap[currentNum] = true
			sum += currentNum
		}
	}

	return sum
}

// This would be far more effective (O(n) instead of O(n^2))
// I got stuck figuring out how to not add the common multiples more than once.
// func SumMultiples(limit int, divisors []int) int {

// 	resultMap := make(map[int]bool)
// 	sum := 0

// 	for _, divisor := range divisors {

// 		n := (limit - 1) / divisor
// 		multiplier := n * (n + 1) / 2
// 		result := multiplier * divisor

// 		if resultMap[result] {
// 			continue
// 		}
// 		fmt.Println(result)

// 		sum += result
// 		resultMap[result] = true
// 	}
// 	return sum
// }
