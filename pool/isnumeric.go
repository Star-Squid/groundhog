package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)

	for _, element := range runes {
		if element >= 48 && element <= 57 {
			continue
		} else {
			return false
		}
	}

	return true
}
