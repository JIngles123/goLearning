package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64=3.4
	v:=reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(7.1)
	// panic: reflect.Value.SetFloat using unaddressable value
	// 因为 v 是不可寻址的，所以不能直接修改它的值
	// 需要先获取它的地址
	
	fmt.Println("settability of v:", v.CanSet())
	v=reflect.ValueOf(&x) // 获取 x 的地址
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	// Elem() 相当于反射里的 *：得到指针指向的那个 float64 对应的 Value。
	// 因为这个值是通过 &x 拿到的，底层变量可寻址，所以再赋值后，可以修改 x 的值。
	// 即，把「指向 x 的指针」变成「可设置的 x 本身」
	v=v.Elem() // 从指针的 reflect.Value 解引用到它指向的那个值。
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // 需要先获取到 x 的地址，才能修改 x 的值
	fmt.Println(v.Interface())
	fmt.Println(v)
}
