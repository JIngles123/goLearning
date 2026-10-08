package main

import (
	"fmt"
)

// 定义一个结构体 stockPosition 及其方法 getValue()
type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

// 定义一个结构体 car 及其方法 getValue()
type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

// 定义一个使用 valuable 接口的函数 showValue(),所有实现了valuable接口的类型都可以调用这个函数
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}
