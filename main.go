package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {

// 	var cat, dog string
// 	cat = "Kalmar"
// 	dog = "Burek"
// 	statement := fmt.Sprintf("I like %v more than %v. \n", cat, dog)

// 	fmt.Println("Hi there!")
// 	fmt.Printf("Are you a %s or a %v person?\n", cat, dog)
// 	fmt.Print(statement)
// }

func summonCat() {
	fmt.Println("- Here, kitty kitty!")
}

func addFifty(x *int) {
	*x += 20
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var name string
	randomAge := rand.Intn(80)
	var age int
	//num := 5

	addFifty(&age)

	summonCat()
	fmt.Println("- What is your name?")
	fmt.Scan(&name)
	fmt.Printf("- Hello, %v!\n- How old are you?\n", name)
	fmt.Scan(&age)
	fmt.Printf("- Wow, %v, you don't look %d!\n", name, age)
	fmt.Printf("- Are you sure you're not %v?\n", randomAge)
}
