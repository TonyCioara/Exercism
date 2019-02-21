package secret

// Handshake creates a secret handshake given the input
func Handshake(input uint) []string {

	factors := []uint{16, 8, 4, 2, 1}

	operations := make(map[uint]bool)

	for _, factor := range factors {
		if input >= factor {
			input -= factor
			operations[factor] = true
		}
	}

	var finalSlice []string

	if operations[1] {
		finalSlice = append(finalSlice, "wink")
	}
	if operations[2] {
		finalSlice = append(finalSlice, "double blink")
	}
	if operations[4] {
		finalSlice = append(finalSlice, "close your eyes")
	}
	if operations[8] {
		finalSlice = append(finalSlice, "jump")
	}
	if operations[16] {
		for left, right := 0, len(finalSlice)-1; left < right; left, right = left+1, right-1 {
			finalSlice[left], finalSlice[right] = finalSlice[right], finalSlice[left]
		}
	}

	return finalSlice
}
