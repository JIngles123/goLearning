package main

import (
	"fmt"
	"time"
)

func main() {
	ch1:=make(chan int)
	ch2:=make(chan int)

	go pump1_151(ch1)
	go pump2_151(ch2)
	go suck_151(ch1, ch2)

	time.Sleep(1e9)
}

func pump1_151(ch chan int) {
	for i:=0;;i++ {
		ch<-i*2
	}
}

func pump2_151(ch chan int) {
	for i:=0;;i++ {
		ch<-i+5
	}
}

func suck_151(ch1, ch2 chan int) {
	for {
		select { // 从不同的并发执行的协程中获取值可以通过关键字 select 去实现
		case v:=<-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v:=<-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

