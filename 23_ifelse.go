package main

import "fmt"

func main() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is greater than 0 and less than 5\n")
	} else {
		fmt.Printf("first is greater than or equal to 5\n")
	}

	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is less than or equal to 10\n")
	}
}

