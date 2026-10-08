package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1:=new(Square)
	sq1.side=5

	var areaIntf Shaper
	// 可以将一个 Square 类型的变量赋值给一个接口类型的变量：areaIntf = sq1
	// 接口变量包含一个指向 Square 变量的引用
	areaIntf=sq1

	// or
	// areaIntf:=Shaper(sq1)
	// or
	// areaIntf:=sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

