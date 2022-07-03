package piscine

func BasicAtoi(s string) int {
	var result int

	// convert string to slice of ascii values
	byteArray := []byte(s)
	intArray := []int{}

	// if between 48-57, subtract 48 (and convert?) to get int
	for i := 0; i < len(byteArray); i++ {
		switch byteArray[i] {
		case 48:
			intArray = append(intArray, 0)
		case 49:
			intArray = append(intArray, 1)
		case 50:
			intArray = append(intArray, 2)
		case 51:
			intArray = append(intArray, 3)
		case 52:
			intArray = append(intArray, 4)
		case 53:
			intArray = append(intArray, 5)
		case 54:
			intArray = append(intArray, 6)
		case 55:
			intArray = append(intArray, 7)
		case 56:
			intArray = append(intArray, 8)
		case 57:
			intArray = append(intArray, 9)
		}
	}

	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		result += intArray[i] * op
		op *= 10
	}

	// if empty return unchanged result
	// if not empty return the int

	return result
}
