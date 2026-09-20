package main

import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
	return
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}

// 输出：
// In function1 at the top
// In function1 at the bottom!
// Function2: Deferred until the end of the calling function!

// 可以看到，function2 函数在 function1 函数结束时被调用，而不是在 function1 函数结束前调用。
// 这是因为 defer 语句会将 function2 函数延迟到 function1 函数结束时再调用。
// 这就是 defer 语句的延迟执行特性。

// 用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块, 一般用于释放某些已分配的资源
// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
// 例如：
// func function1() {
// 	defer function2()
// 	defer function3()
// 	defer function4()
// }
// 执行顺序为：function4 -> function3 -> function2 -> function1
// 这是因为 defer 语句会将 function4 函数延迟到 function1 函数结束时再调用，
// 然后将 function3 函数延迟到 function4 函数结束时再调用，
// 然后将 function2 函数延迟到 function3 函数结束时再调用，
// 最后将 function1 函数延迟到 function2 函数结束时再调用。