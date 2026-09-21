package main

import "fmt"

func f(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	var ar [3]int
	f(ar) // 值传递
	fp(&ar) // 指针传递
}

// 输出
// [0 0 0]
// &[0 0 0]

// fmt 包对指针有特殊处理，它会「智能地」帮你解引用，除非你明确要求打印地址
// fmt 在遇到指针时（用 %v / Println 等默认格式）会这样处理：
//            情况	                         输出
// 指向结构体、数组、切片、map 等的指针	     & + 它指向的值的格式
// 指向基础类型（int、string 等）的指针	    直接打印地址
// 用 %p	                            一律打印地址
// 用 %v 但指针指向基础类型	               打印地址