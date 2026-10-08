package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1:=new(Square)
	sq1.side=5

	areaIntf=sq1
	// 固定写法，判断 areaIntf 里是否包含一个 'Square' 类型的变量
	// 一定要是指针类型，因为 Area() 方法的接收者是 *Square
	if t,ok:=areaIntf.(*Square); ok{
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if u,ok:=areaIntf.(*Circle); ok{
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
