package main

import "fmt"

func main() {
	addressBook := map[string]string{
		"Jannet": "22 Water St",
		"Joe":    "241 North Rd",
		"Robert": "86 Stone St",
	}

	for name := range addressBook {
		fmt.Println(name, "lives somewhere")
	}

	for name, address := range addressBook {
		fmt.Println(name, "lives at", address)
	}

	nums := [4]int{1, 2, 3, 4}
	//secret := 0

	for i := range nums {
		fmt.Println(i)
	}

	//fmt.Println(nums)
}
