package piscine

func AppendRange(min, max int) []int {
	result := []int{}

	if min >= max {
		return []int(nil)
	} else {
		for i := min; i < max; i++ {
			result = append(result, i)
		}
	}

	return result
}
