package piscine

func IsLower(s string) bool {
	runes := []rune(s)

	for _, element := range runes {
		if element >= 97 && element <= 122 {
			continue
		} else {
			return false
		}
	}

	return true
}
