package main

import "fmt"

// 一个结构体上的简单方法示例

type TwoInts struct {
	a int
	b int
}

func main() {
	two1:=new(TwoInts)
	two1.a=12
	two1.b=10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2:=TwoInts{3,4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a+tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a+tn.b+param
}

// 定义方法
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

// Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。

