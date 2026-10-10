package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

/*
panic 可以直接从代码初始化：当错误条件（所测试的代码）很严苛且不可恢复，程
序不能继续运行时，可以使用 panic 函数产生一个中止程序的运行时错误。panic 接收
一个做任意类型的参数，通常是字符串，在程序死亡时被打印出来。Go 运行时负责中止程序
并给出调试信息。

在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语
句都会保证执行并把控制权交还给接收到 panic 的函数调用者。这样向上冒泡直到最顶层，
并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情
况：这个终止过程就是 panicking。

*/