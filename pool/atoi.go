package piscine

func Atoi(s string) int {
	var result int

	var isInvalid bool
	var isNegative bool

	byteArray := []byte(s)
	intArray := []int{}

	// it is only negative if ONE minus at the front
	// one + at the front is positive

	// + or - in other positions than [0] means invalid

	if len(s) < 1 {
		isInvalid = true
	} else {

		for i := 0; i < len(byteArray); i++ {
			switch byteArray[i] {
			case 45: // minus sign
				intArray = append(intArray, 0) // so the array doesn't lose length
			case 43: // plus sign
				intArray = append(intArray, 0) // so the array doesn't lose length
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
			default:
				isInvalid = true
			}
		}

		// test if negative or invalid now
		if byteArray[0] == 45 {
			isNegative = true
		}

		if len(byteArray) > 1 {
			for i := 1; i < len(byteArray); i++ {
				if byteArray[i] == 45 || byteArray[i] == 43 {
					isInvalid = true
				}
			}
		}
	}

	if isInvalid {
		result = 0
	} else {
		op := 1
		for i := len(s) - 1; i >= 0; i-- {
			result += intArray[i] * op
			op *= 10
		}
	}

	if isNegative {
		result = 0 - result
	}

	return result
}
