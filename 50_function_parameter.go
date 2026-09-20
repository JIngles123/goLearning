package main

import (
	"fmt"
)

func main() {
	// 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
