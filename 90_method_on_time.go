package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time // 匿名内嵌
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m:=myTime{time.Now()}
	// 调用匿名内嵌类型的方法
	fmt.Println("Full time now: ", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars: ", m.first3Chars())
}

/*
函数将变量作为参数：Function1(recv)
方法在变量上被调用：recv.Method1()

在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，
即通过引用调用时，函数也可以改变参数的状态）。
不要忘记 Method1 后边的括号 ()，否则会引发编译器错误
*/