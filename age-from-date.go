package age

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Basic ask and print
// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("Type here:")
// 	scanner.Scan()
// 	input := scanner.Text()
// 	fmt.Printf("%q is a naughty word!", input)
// }

func age() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your birth year here:")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("You are %d years old", 2022-input)
}
