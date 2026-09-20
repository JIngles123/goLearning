package main

import "fmt"

func main() {
	var i int = 5

	// 基于条件判断的迭代
	for i >= 0 {
		i=i-1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
