package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Println(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x+=delta
		return x
	}
}

// 三次调用函数 f 的过程中函数 Adder () 中变量 delta 的值分别为：1、20 和 300。
// 我们可以看到，在多次调用中，变量 x 的值是被保留的
// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
