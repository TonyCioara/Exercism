package grains

import (
	"errors"
	"math"
)

// Square returns the amount of grains on a given square of the chess board
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		err := errors.New("Invalid input")
		return 0, err
	}

	return uint64(math.Pow(2, float64(input-1))), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		result, _ := Square(i)
		sum += result
	}
	return sum
}
