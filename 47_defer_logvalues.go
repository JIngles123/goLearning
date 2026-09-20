package main

import (
	"io"
	"log"
)

// defer   func() { ... }   ()
//  │         │              │
//  │         │              └── ④ 立即调用这个函数
//  │         └── ② 一个匿名函数（函数值）
//  └── ① defer 关键字
// %q —— 带引号的字符串
// %d —— 十进制整数
// %v —— 万能默认格式

func func1(s string) (n int, err error) {
	// 匿名函数在函数返回时被调用
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()

	return 7, io.EOF
}

func main() {
	func1("Go")
}

// 输出：
// 2026/09/20 19:41:24 func1("Go") = 7, EOF