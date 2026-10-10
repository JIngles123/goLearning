package main

import (
	"fmt"
	"time"
)

func main() {
	// 通道也是引用类型，所以我们使用 make() 函数来给它分配内存。这里先声明了一个字符串通道 ch，然后创建了它（实例化）
	ch:=make(chan string)

	// 启动两个协程，一个用于发送数据，一个用于接收数据
	// 如果两个协程需要通信，必须给它们同一个通道作为参数
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9) // 给程序1秒的时间来运行两个协程
}
/*
协程之间的同步非常重要：
1、main () 等待了 1 秒让两个协程完成，如果不这样，sendData () 就没有机会输出。
2、getData () 使用了无限循环：它随着 sendData () 的发送完成和 ch 变空也结束了。
3、如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 panic
*/

func sendData(ch chan string) { // 发送5个字符串到通道 ch，其中 <-表示通信操作符
	ch<-"Washington"
	ch<-"Tripoli"
	ch<-"London"
	ch<-"Beijing"
	ch<-"Tokio"
}

func getData(ch chan string) { // 按顺序接收内容并打印
	var input string
	for {
		input = <-ch // 从通道里取出一个值的表达式
		fmt.Printf("%s ", input)
	}
}
// <- 表示数据往哪边流动，ch <- xxx 表示数据流向通道 ch，xxx = <- ch 表示从通道 ch 中接收数据
