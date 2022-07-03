package piscine

func SortIntegerTable(table []int) {
	swap := func(a *int, b *int) {
		if *a > *b {
			var c int = 3
			c = *b
			// c holds the value of b, so we can assign b=value of a, a = value of c (originally b)

			*b = *a
			*a = c
		}
	}

	// for every element, the loop runs through the whole slice
	for i := 0; i < len(table); i++ {
		for i := 0; i < len(table)-1; i++ {
			swap(&table[i], &table[i+1])
		}
	}
}
