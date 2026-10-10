package main

import (
	"errors"
	"fmt"
)

// 任何时候当你需要一个新的错误类型，都可以用 errors（必须先 import）包
// 的 errors.New 函数接收合适的错误信息来创建
var errNotFound error = errors.New("Not found error")

func main() {
	// 由于 fmt.Printf 会自动调用 String() 方法，所以错误
	// 信息 “Error: math - square root of negative number” 会打印出来。
	fmt.Printf("error: %v", errNotFound)
}
