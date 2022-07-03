package piscine

func AlphaCount(s string) int {
	runes := []rune(s)
	validLetters := 0

	for i := 0; i <= len(runes)-1; i++ {
		if (runes[i] >= 65 && runes[i] <= 90) || (runes[i] >= 97 && runes[i] <= 122) {
			validLetters++
		}
	}

	return validLetters
}
