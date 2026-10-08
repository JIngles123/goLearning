package main

import "fmt"

var i=5
var str="ABC"

type Person struct {
	name string
	age int
}

type Any interface{} // 空接口或者最小接口 不包含任何方法，它对实现不做任何要求

func main() {
	var val Any // 可以给一个空接口类型的变量 var val interface {} 赋任何类型的值。
	val=5
	fmt.Printf("val has the value: %v\n", val)
	val=str
	fmt.Printf("val has the value: %v\n", val)
	
	pers1:=new(Person)
	pers1.name="Rob Pike"
	pers1.age=55
	val=pers1
	fmt.Printf("val has the value: %v\n", val)

	switch t:=val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
