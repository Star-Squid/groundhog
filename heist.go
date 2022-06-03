package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards.")
	} else {
		isHeistOn = false
		fmt.Println("The guards spotted you. Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 30 {
		fmt.Println("You're in the vault. Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault can't be opened")
	} else {
		fmt.Println("You fully failed!")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Caught on camera - busted!")
		case 1:
			isHeistOn = false
			fmt.Println("Tripped the alarm - busted!")
		case 2:
			isHeistOn = false
			fmt.Println("You got cold feet and returned the money, you loser.")
		case 4:
			isHeistOn = false
			fmt.Println("A guard got hold of you - busted!")
		default:
			isHeistOn = true
			fmt.Println("You made it to the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("You've stolen %v!\n", amtStolen)
	}

	fmt.Println("Is the heist still on?", isHeistOn)
}
