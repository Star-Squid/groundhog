package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)

	for _, element := range runes {
		if element == 32 || element >= 65 && element <= 90 || element >= 97 && element <= 122 || element >= 48 && element <= 57 {
			continue
		} else {
			return false
		}
	}

	return true
}
