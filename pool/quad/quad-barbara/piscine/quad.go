package piscine

import "fmt"

func QuadA(x, y int) {

	//1. return nothing if arguments are not positive non-zero integers
	if x <= 0 || y <= 0 {
		return
	}

	//2. define all building blocks we will need - this is repetitive here but easily swapped & descriptive
	blocks := map[string]string{
		"topLeftCorner":  "o",
		"horizontalLine": "-",
		"topRightCorner": "o",

		"verticalLine": "|",
		"empty":        " ",

		"bottomLeftCorner":  "o",
		"bottomRightCorner": "o",
	}

	//3. define actions needed to create each line
	lineOne := func() {
		//top left corner prints every time
		fmt.Print(blocks["topLeftCorner"])

		//top line only prints if x is 3 or more
		if x >= 3 {
			//print it as many times as x-2
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		//top right corner only prints if x is 2 or over
		if x >= 2 {
			fmt.Print(blocks["topRightCorner"])
		}

		//start the next line
		fmt.Print("\n")
	}

	lineMiddle := func() {

		fmt.Print(blocks["verticalLine"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["empty"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["verticalLine"])
		}

		fmt.Print("\n")
	}

	lineBottom := func() {

		fmt.Print(blocks["bottomLeftCorner"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["bottomRightCorner"])
		}

		fmt.Print("\n")
	}

	//4. build the rectangle:
	//we always need the first line

	lineOne()

	//we need the middle line(s) if y is 3 or larger
	if y >= 3 {

		//print middle line as many times as y-2
		for i := y - 2; i > 0; i-- {
			lineMiddle()
		}

	}

	//only print the last line if y is 2 or larger
	if y >= 2 {
		lineBottom()
	}

}

func QuadB(x, y int) {

	//1. return nothing if arguments are not positive non-zero integers
	if x <= 0 || y <= 0 {
		return
	}

	//2. define all building blocks we will need - this is repetitive but easily swapped & descriptive
	blocks := map[string]string{
		"topLeftCorner":  "/",
		"horizontalLine": "*",
		"topRightCorner": "\\",

		"verticalLine": "*",
		"empty":        " ",

		"bottomLeftCorner":  "\\",
		"bottomRightCorner": "/",
	}

	//3. define actions needed to create each line
	lineOne := func() {
		//top left corner prints every time
		fmt.Print(blocks["topLeftCorner"])

		//top line only prints if x is 3 or more
		if x >= 3 {
			//print it as many times as x-2
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		//top right corner only prints if x is 2 or over
		if x >= 2 {
			fmt.Print(blocks["topRightCorner"])
		}

		//start the next line
		fmt.Print("\n")
	}

	lineMiddle := func() {

		fmt.Print(blocks["verticalLine"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["empty"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["verticalLine"])
		}

		fmt.Print("\n")
	}

	lineBottom := func() {

		fmt.Print(blocks["bottomLeftCorner"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["bottomRightCorner"])
		}

		fmt.Print("\n")
	}

	//4. build the rectangle:
	//we always need the first line

	lineOne()

	//we need the middle line(s) if y is 3 or larger
	if y >= 3 {

		//print middle line as many times as y-2
		for i := y - 2; i > 0; i-- {
			lineMiddle()
		}

	}

	//only print the last line if y is 2 or larger
	if y >= 2 {
		lineBottom()
	}

}

func QuadC(x, y int) {

	//1. return nothing if arguments are not positive non-zero integers
	if x <= 0 || y <= 0 {
		return
	}

	//2. define all building blocks we will need - this is repetitive but easily swapped & descriptive
	blocks := map[string]string{
		"topLeftCorner":  "A",
		"horizontalLine": "B",
		"topRightCorner": "A",

		"verticalLine": "B",
		"empty":        " ",

		"bottomLeftCorner":  "C",
		"bottomRightCorner": "C",
	}

	//3. define actions needed to create each line
	lineOne := func() {
		//top left corner prints every time
		fmt.Print(blocks["topLeftCorner"])

		//top line only prints if x is 3 or more
		if x >= 3 {
			//print it as many times as x-2
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		//top right corner only prints if x is 2 or over
		if x >= 2 {
			fmt.Print(blocks["topRightCorner"])
		}

		//start the next line
		fmt.Print("\n")
	}

	lineMiddle := func() {

		fmt.Print(blocks["verticalLine"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["empty"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["verticalLine"])
		}

		fmt.Print("\n")
	}

	lineBottom := func() {

		fmt.Print(blocks["bottomLeftCorner"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["bottomRightCorner"])
		}

		fmt.Print("\n")
	}

	//4. build the rectangle:
	//we always need the first line

	lineOne()

	//we need the middle line(s) if y is 3 or larger
	if y >= 3 {

		//print middle line as many times as y-2
		for i := y - 2; i > 0; i-- {
			lineMiddle()
		}

	}

	//only print the last line if y is 2 or larger
	if y >= 2 {
		lineBottom()
	}

}

func QuadD(x, y int) {

	//1. return nothing if arguments are not positive non-zero integers
	if x <= 0 || y <= 0 {
		return
	}

	//2. define all building blocks we will need - this is repetitive but easily swapped & descriptive
	blocks := map[string]string{
		"topLeftCorner":  "A",
		"horizontalLine": "B",
		"topRightCorner": "C",

		"verticalLine": "B",
		"empty":        " ",

		"bottomLeftCorner":  "A",
		"bottomRightCorner": "C",
	}

	//3. define actions needed to create each line
	lineOne := func() {
		//top left corner prints every time
		fmt.Print(blocks["topLeftCorner"])

		//top line only prints if x is 3 or more
		if x >= 3 {
			//print it as many times as x-2
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		//top right corner only prints if x is 2 or over
		if x >= 2 {
			fmt.Print(blocks["topRightCorner"])
		}

		//start the next line
		fmt.Print("\n")
	}

	lineMiddle := func() {

		fmt.Print(blocks["verticalLine"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["empty"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["verticalLine"])
		}

		fmt.Print("\n")
	}

	lineBottom := func() {

		fmt.Print(blocks["bottomLeftCorner"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["bottomRightCorner"])
		}

		fmt.Print("\n")
	}

	//4. build the rectangle:
	//we always need the first line

	lineOne()

	//we need the middle line(s) if y is 3 or larger
	if y >= 3 {

		//print middle line as many times as y-2
		for i := y - 2; i > 0; i-- {
			lineMiddle()
		}

	}

	//only print the last line if y is 2 or larger
	if y >= 2 {
		lineBottom()
	}

}

func QuadE(x, y int) {

	//1. return nothing if arguments are not positive non-zero integers
	if x <= 0 || y <= 0 {
		return
	}

	//2. define all building blocks we will need - this is repetitive but easily swapped & descriptive
	blocks := map[string]string{
		"topLeftCorner":  "A",
		"horizontalLine": "B",
		"topRightCorner": "C",

		"verticalLine": "B",
		"empty":        " ",

		"bottomLeftCorner":  "C",
		"bottomRightCorner": "A",
	}

	//3. define actions needed to create each line
	lineOne := func() {
		//top left corner prints every time
		fmt.Print(blocks["topLeftCorner"])

		//top line only prints if x is 3 or more
		if x >= 3 {
			//print it as many times as x-2
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		//top right corner only prints if x is 2 or over
		if x >= 2 {
			fmt.Print(blocks["topRightCorner"])
		}

		//start the next line
		fmt.Print("\n")
	}

	lineMiddle := func() {

		fmt.Print(blocks["verticalLine"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["empty"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["verticalLine"])
		}

		fmt.Print("\n")
	}

	lineBottom := func() {

		fmt.Print(blocks["bottomLeftCorner"])

		if x >= 3 {
			for i := x - 2; i > 0; i-- {
				fmt.Print(blocks["horizontalLine"])
			}
		}

		if x >= 2 {
			fmt.Print(blocks["bottomRightCorner"])
		}

		fmt.Print("\n")
	}

	//4. build the rectangle:
	//we always need the first line

	lineOne()

	//we need the middle line(s) if y is 3 or larger
	if y >= 3 {

		//print middle line as many times as y-2
		for i := y - 2; i > 0; i-- {
			lineMiddle()
		}

	}

	//only print the last line if y is 2 or larger
	if y >= 2 {
		lineBottom()
	}

}
