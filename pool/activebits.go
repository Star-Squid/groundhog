package piscine

func ActiveBits(n int) int {
	// convert to byte
	// convert byte to array of 0 and 1
	// loop over array counting 1s
	var result int = n
	if n <= 255 {
		result = countBits(n)
	}
	return result
}

func countBits(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n & 1) + countBits(n>>1)
	}
}

// but this is not really a byte, but an int that represents a byte
// so instead:

func convertIntToArray(n int) []int {
	result := []int{}

	if n != 0 {
		i := n % 10
		result = append(result, i)
		convertIntToArray(n / 10)
	}
	return result
}

// do we need this? no

func convertDecimalToBinary(n int) int {
	binary := 0
	counter := 1
	remainder := 0

	for n != 0 {
		remainder = n % 2
		n = n / 2
		binary += remainder * counter
		counter *= 10

	}
	return binary
}
