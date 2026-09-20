package main

import "fmt"

func main() {
	f()
}

func f() {
	for i:=0;i<4;i++ {
		// 定义一个匿名函数
		g:=func(i int) {fmt.Printf("%d ", i)}
		// 调用匿名函数
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

