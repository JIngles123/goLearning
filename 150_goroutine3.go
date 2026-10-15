package main

import "fmt"

func main() {
	ch:=make(chan string)
	go sendData_150(ch) // 启动一个goroutine来发送数据,协程
	getData_150(ch) // getData_150和main在同一个线程里
}

func sendData_150(ch chan string) {
	ch<-"Washington"
	ch<-"Tripoli"
	ch<-"London"
	ch<-"Beijing"
	ch<-"Tokyo"
	// 只有发送者需要关闭通道，接收者永远不会需要
	close(ch) //关闭通道，不需要每次都关闭，只有当需要告诉接收者不会再提供新的值的时候，才需要关闭
}

func getData_150(ch chan string) {
	for {
		input,open:=<-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

