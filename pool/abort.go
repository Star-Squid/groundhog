package piscine

func Abort(a, b, c, d, e int) int {
	// sort the numbers
	// c is the median

	table := []int{a, b, c, d, e}
	SortIntegerTable(table)
	return table[2]
}
