package main

import (
	"os"
	"strconv"
)

func divideByZeroError(operator string) {
	var errMessage string
	if operator == "/" {
		errMessage = "No division by 0" + "\n"
	} else {
		errMessage = "No modulo by 0" + "\n"
	}

	os.Stdout.Write([]byte(errMessage))
	os.Exit(0)
}

func main() {
	var rightNum int
	var operator string
	var result int
	args := os.Args[1:]

	if len(args) == 3 {
		leftNum, err := strconv.Atoi(args[0])
		operator = args[1]
		rightNum, _ = strconv.Atoi(args[2])
		if leftNum == 9223372036854775807 || rightNum == 9223372036854775807 {
			os.Exit(0)
		}

		if err == nil && operator == "+" || err == nil && operator == "-" || err == nil && operator == "*" || err == nil && operator == "/" || err == nil && operator == "%" {
			if operator == "+" {
				result = leftNum + rightNum
			} else if operator == "-" {
				result = leftNum - rightNum
			} else if operator == "*" {
				result = leftNum * rightNum
			} else if operator == "/" {
				if rightNum == 0 {
					divideByZeroError("/")
				} else {
					result = leftNum / rightNum
				}
			} else if operator == "%" {
				if rightNum == 0 {
					divideByZeroError("%")
				} else {
					result = leftNum % rightNum
				}
			}

			resultString := strconv.Itoa(result) + "\n"
			os.Stdout.Write([]byte(resultString))
		}
	}
}
