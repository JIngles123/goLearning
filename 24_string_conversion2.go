package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	
	an, err := strconv.Atoi(orig)
	if err != nil { // 判断err是否包含一个真正的错误
		fmt.Printf("orig %s is not an integer - exiting with error, an is %d\n", orig, an)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}

// 示例1：常用写法
// value, err := pack1.Function1(param1)
// if err != nil {
//     fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
//     return err  退出程序的话使用语句：os.Exit(1)
// }
// // 未发生错误，继续执行

// 示例2：打开一个名为name的只读文件，并返回一个文件对象
// f, err := os.Open(name)
// if err != nil {
//     return err
// }
// doSomething(f) // 当没有错误发生时，文件对象被传入到某个函数中
// doSomething
