package main

import "fmt"

func BasicAtoi(s string) rune { //NEEDS TO RTRN INT
	//if s can be converted to int, return that, if not, return 0
	//f the string is not considered as a valid number
	//transforms a number defined as a string in a number defined as an int
	bytes := []byte(s)
	num := int(bytes[0]) //gives ascii value
	character := rune(num)

	if character == '1' || character == '2' || character == '3' || character == '4' || character == '5' || character == '6' || character == '7' || character == '8' || character == '9' {
		return character
	} else {
		return '0'
	}

}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
	fmt.Println(BasicAtoi("hello"))

}
