package piscine

func Any(f func(string) bool, a []string) bool {
	var isTrue bool
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			isTrue = true
		}
	}
	return isTrue
}
