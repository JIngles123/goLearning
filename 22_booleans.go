package main

import "fmt"

func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value of bool1 is true\n")
	} else {
		fmt.Printf("The value of bool1 is false\n")
	}

	bool2 := (1 == 2)
	if bool2 {
		fmt.Printf("The value of (1 == 2) is true\n")
	} else {
		fmt.Printf("The value of (1 == 2) is false\n")
	}
}