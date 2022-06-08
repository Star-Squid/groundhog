package main

import (
	"fmt"
)

func main() {
	//goal: ask for password to get in, lock out after 3 tries

	//make var for correct passwird, string
	var correct string
	correct = "please"

	//make var int for number of tries
	var numOfTries int
	numOfTries = 0

	for numOfTries = 1; numOfTries <= 3; numOfTries++ {
		var guess string
		fmt.Println("What's the password?")
		fmt.Scan(&guess)

		if guess == correct {
			fmt.Println("You're in!")
			break
		} else {
			triesRemaining := 3 - numOfTries
			fmt.Println("Incorrect password! Tries remaining:", triesRemaining)
			if numOfTries == 3 {
				fmt.Println("This isn't working. Goodbye.")
			}
		}
	}
}
