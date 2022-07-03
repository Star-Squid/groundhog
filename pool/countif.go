package piscine

func CountIf(f func(string) bool, tab []string) int {
	var isTrue int
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			isTrue++
		}
	}
	return isTrue
}
