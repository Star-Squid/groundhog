package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	} else {
		j := 0
		result := make([]int, max-min)
		for i := min; i < max; i++ {
			result[j] = i
			j++
		}
		return result
	}
}
