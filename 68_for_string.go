package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
}
/*
Go 语言中的字符串是不可变的，也就是说 str[index] 这样的表达式是不可以被放在等号左侧的。
如果尝试运行 str[i] = 'D' 会得到错误：cannot assign to str[i]。

因此，必须先将字符串转换成字节数组，然后再通过修改数组中的元素值来达到修改字符串的目的，
最后将字节数组转换回字符串格式。

如：
s := "hello"
c := []byte(s)
c[0] = 'c'
s2 := string(c) // s2 == "cello"

*/