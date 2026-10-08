package main

import "fmt"

type specialString string
var whatIsThis specialString="hello"

func TypeSwitch() {
	// 空接口在 type-switch 中联合匿名函数 (lambda 函数) 的用法
	testFunc:=func(any interface{}) {
		switch v:=any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type")
		}
	}
	testFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}
