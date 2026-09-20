package main

import "fmt"

func main() {
	var num1 int = 7

	switch {
	case num1 < 0:
		fmt.Println("Number is negative"); fallthrough
	case num1>0 && num1<10:
		fmt.Println("Number is between 0 and 10"); fallthrough // 继续执行下一个case
	default:
		fmt.Println("Number is 10 or greater")
	}
}

