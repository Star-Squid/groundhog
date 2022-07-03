package piscine

func Max(a []int) int {
	var result int

	if len(a) <= 0 {
		return 0
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] > result {
				result = a[i]
			}
		}
	}

	return result
}
