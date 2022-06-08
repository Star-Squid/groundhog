package main

import (
	"fmt"
)

func main() {
	jellybeans := []string{"green", "blue", "yellow", "red", "green", "yellow", "red"}
	for index := 0; index < len(jellybeans); index++ {
		if jellybeans[index] == "green" {
			continue
		}
		fmt.Println("You ate the", jellybeans[index], "jellybean!")
	}
}
