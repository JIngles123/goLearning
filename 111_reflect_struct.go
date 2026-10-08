package main

import (
	"fmt"
	"reflect"
)

type NotKnownType struct {
	s1, s2, s3 string
}

func (n NotKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotKnownType{"Ada", "Go", "Oberon"}

func main() {
	value:=reflect.ValueOf(secret)
	typ:=reflect.TypeOf(secret)// 通过反射获取 secret 的类型，这里的 typ 是 *NotKnownType
	fmt.Println(typ)
	knd:=value.Kind() // 获取 value 的底层类型
	fmt.Println(knd)

	for i:=0;i<value.NumField();i++ { // NumField()返回结构体内的字段数量
		fmt.Printf("Field %d: %v\n", i, value.Field(i)) // Field(i)返回第 i 个字段的值
	}

	results:=value.Method(0).Call(nil) // 用反射调用结构体上的第 0 个方法，并把返回值收进 results
	// 或者按名字调用方法：value.MethodByName("String").Call(nil)
	fmt.Println(results)
}

