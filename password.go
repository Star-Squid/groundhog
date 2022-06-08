package main

import (
	"fmt"
)

func main() {
	correct := "please"
	tries := 0
	keepGoing := true

	// if tries == 3 {
	// 	keepGoing = false
	// }

	for tries = 0; keepGoing; tries++ {
		var guess string
		fmt.Println("- What's the magic word?")
		fmt.Scan(&guess)
		if guess == correct {
			fmt.Println("- You're in!")
			keepGoing = false
		} else {
			fmt.Println("- Password incorrect")
			if tries >= 2 {
				fmt.Println("- Too many tries. You lost your chance.")
				keepGoing = false
			}
		}
	}

}
