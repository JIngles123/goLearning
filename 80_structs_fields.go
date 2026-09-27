package main

import "fmt"

/*
结构体定义：
type T struct {a, b int}
T 结构体类型有两个字段，a 和 b，都是 int 类型。
*/

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1) // 使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	/*
	   定义1、
	        ms := &struct1{10, 15.5, "Chris"}
	        // 此时ms的类型是 *struct1，底层仍然会调用 new(struct1)

	   定义2、
	        var ms struct1
	        ms = struct1{10, 15.5, "Chris"}
	*/

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)

	fmt.Println(ms)
}
