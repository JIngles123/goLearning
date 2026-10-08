package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t:=T{23, "skidoo"}
	s:=reflect.ValueOf(&t).Elem() // 通过反射获取 t 的地址，并得到一个反射对象,此时s可寻址，可以修改t的值
	typeOfT:=s.Type()
	
	for i:=0;i<s.NumField();i++ { // NumField()返回结构体内的字段数量
		f:=s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
		typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	// 通过反射修改结构体的字段值,只有被导出的字段才能被修改
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
