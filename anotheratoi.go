package main

import "fmt"

func BasicAtoi(s string) int { //NEEDS TO RTRN INT
	//if s can be converted to int, return that, if not, return 0
	//f the string is not considered as a valid number
	//transforms a number defined as a string in a number defined as an int

	runes := []rune(s)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	//result for golang [71 79 76 65 78 71] //array of ascii values//since it's a rune, maybe unicode
	//var hm = new(*int)

	for j := 0; j < len(result); j++ {

		switch result[j] {
		case 49:
			result[j] = 1
		case 50:
			result[j] = 2
		case 51:
			result[j] = 3
		case 52:
			result[j] = 4
		case 53:
			result[j] = 5
		case 54:
			result[j] = 6
		case 55:
			result[j] = 7
		case 56:
			result[j] = 8
		case 57:
			result[j] = 9
		default:
			continue
		}
	}

	var conversion int

	for k := 0; k < len(result); k++ {
		if k > 0 {
			conversion += result[k] + k*10
		} else {
			conversion += result[k]
		}
	}
	return conversion
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi("hello"))

}
