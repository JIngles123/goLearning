package main

import (
	"fmt"
	"time"
)

/*
从指定通道中读取数据知道通道关闭，才继续执行下边的代码
另一个协程必须写入ch，否则会阻塞，而且必须在写入完成后才能关闭
*/

func main() {
	suck_147(pump_147())
	time.Sleep(1e9) // 1e9 = 10^9 = 1秒
}

func pump_147() chan int {
	ch:=make(chan int)
	go func() {
		for i:=0;;i++ {
			ch<-i
		}
	}()
	return ch
}

func suck_147(ch chan int) {
	go func() {
		for v:=range ch {
			fmt.Println(v)
		}
	}()
}
