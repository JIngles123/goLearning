package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan int) // 一个无缓冲的通道，只能传递一个数据
	go pump_145(ch1)
	go suck(ch1)
	time.Sleep(1e9)// 1e9 = 10^9 = 1秒
	fmt.Println(<-ch1,"print1") // 这里接收一次通道数据，并打印，但pump继续发后面的数时，没有别的接受者，发送端就阻塞，即没东西打印
	fmt.Println(<-ch1,"print2")
}

func pump_145(ch chan int) { // 将数据发送到通道ch1 (生产者)
	for i:=0;;i++ {
		ch<-i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
