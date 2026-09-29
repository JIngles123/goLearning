package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1:=new(TwoInts)
	two1.a=12
	two1.b=10

	// 格式化描述符:
	// %T 会给出类型的完全规格
	// %#v 会给出实例的完整输出，包括它的字段
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("tow1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a)+"/"+strconv.Itoa(tn.b)+")"
}
