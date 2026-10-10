package main

import (
	"fmt"
)

/*
这个（recover）内建函数被用于从 panic 或 错误场景中恢复：让程序可以
从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。

recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来
的错误值，如果是正常执行，调用 recover 会返回 nil，且没有其它效果。
总结：panic 会导致栈被展开直到 defer 修饰的 recover () 被调用或者程序中止。

这个例子结合了panic defer 和 recover 的用法
*/

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e:=recover(); e!=nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	// 以下内容不print，因为 badCall 中调用了 panic,程序会崩溃，如果defer里
	// 的recover成功，则外层调用方会继续执行，否则会直接退出程序
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

