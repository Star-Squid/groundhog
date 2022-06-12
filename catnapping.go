package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	guests := [5]string{"Mat", "Alastair", "Dr Frankfurter", "Blorbo", "Obama"}
	storage := [6]string{"behind the sofa", "inside the armchair", "under the blanket", "in the cupboard", "in their pocket", "in their backpack"}

	guestSlice := guests[:]
	storageSlice := storage[:]

	guilty := getRandomElement(guestSlice)
	hidingPlace := getRandomElement(storageSlice)

	fmt.Println("The cat was hidden by", guilty, hidingPlace, "!")

}

func getRandomElement(slice []string) string {
	randomIndex := rand.Intn(len(slice))
	// fmt.Println(randomIndex)
	return slice[randomIndex]
}
