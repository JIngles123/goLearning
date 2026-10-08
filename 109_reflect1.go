package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64=3.4
	fmt.Println("type:", reflect.TypeOf(x))

	v:=reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float()) // 返回这个 float64 类型的实际值
	
	fmt.Println(v.Interface())// Interface returns v's current value as an interface{}.
	fmt.Printf("value is %5.2e\n", v.Interface())
	y:=v.Interface().(float64)
	fmt.Println(y)
}

