package main

import (
	"fmt"
	"os"
)

func main() {
	allArgs := os.Args[1:]

	if len(allArgs) == 1 {
		filename := allArgs[0]

		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		arr := make([]byte, 14)

		file.Read(arr)
		fmt.Println(string(arr))

		file.Close()
	} else if len(allArgs) > 1 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("File name missing")
	}
}
