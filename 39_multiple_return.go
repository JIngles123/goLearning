package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// 非命名返回值，return时，将返回所有参数的值
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名返回值
// 命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，
// 当需要返回的时候，我们只需要一条简单的不带参数的 return 语句。
// 需要注意的是，即使只有一个命名返回值，也需要使用 () 括起来
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
	// or return x2, x3
	// or return x3, x2
}
