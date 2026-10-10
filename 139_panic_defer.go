package main

import (
	"fmt"
)

func main() {
	f_139()
	fmt.Println("Returned normally from f_139")
}

/*
defer   func() { ... }   ()
 │         │              │
 │         │              └── ③ 调用这个函数（那对空括号）
 │         └── ② 一个匿名函数（本身只是个值）
 └── ① defer 关键字
*/
func f_139() {
	defer func() {
		if r:=recover(); r!=nil {
			fmt.Println("Recovered in f_139", r)
		}
	}()

	fmt.Println("Calling g_139")
	g_139(0)
	fmt.Println("Returned normally from g_139")
}

func g_139(i int) {
	if i>3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in g_139", i)
	fmt.Println("Printing in g_139", i)
	g_139(i+1)
}

/* 执行结果：
$ go run 139_panic_defer.go 
Calling g_139
Printing in g_139 0
Printing in g_139 1
Printing in g_139 2
Printing in g_139 3
Panicking!
Defer in g_139 3
Defer in g_139 2
Defer in g_139 1
Defer in g_139 0
Recovered in f_139 4
Returned normally from f_139

在递归调用中panic后，panic异常被恢复，defer语句会保证执行，
但是defer语句的执行顺序是后进先出，所以defer语句会先执行defer in g_139 3，
再执行defer in g_139 2，再执行defer in g_139 1，再执行defer in g_139 0，
最后执行Recovered in f_139 4。
*/
