package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
}

// 无法得到一个文字或者常量的地址
// 如：
// const i = 5
// ptr := &i //error: cannot take the address of i
// ptr2 := &10 //error: cannot take the address of 10
// c = *p++  pointer+2 // 不允许，go中的指针保证内存的安全

/*
// 对一个空指针的反向引用是不合法的，并且会使程序崩溃
package main
func main() {
    var p *int = nil
    *p = 0
}
// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
// runtime error: invalid memory address or nil pointer dereference
*/
