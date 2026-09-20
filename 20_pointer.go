package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's address is %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory address %p is %d\n", intP, *intP)
}
