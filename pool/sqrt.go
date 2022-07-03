package piscine

func Sqrt(nb int) int {
	// test if nb is 0 or 1, special cases
	if nb <= 0 {
		return 0
	} else if nb == 1 {
		return 1
	}

	// test every value starting at 1
	i := 1

	for i <= nb {
		if i*i == nb {
			break
		} else if i >= nb { // how did it work when we just had i == nb??
			return 0
		}
		i++
		// we don't test for numbers that don't have a non-int root?
		// i think we do and they return 0
	}
	return i
}
