package main

import "fmt"

// func main() {

// 	var cat, dog string
// 	cat = "Kalmar"
// 	dog = "Burek"
// 	statement := fmt.Sprintf("I like %v more than %v. \n", cat, dog)

// 	fmt.Println("Hi there!")
// 	fmt.Printf("Are you a %s or a %v person?\n", cat, dog)
// 	fmt.Print(statement)
// }

func main() {
	var name string

	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v!\n", name)
}
