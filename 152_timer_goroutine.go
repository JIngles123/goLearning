package main

import (
	"fmt"
	"time"
)

func main() {
	tick:=time.Tick(1e8) // 1e8 = 100000000 = 100ms 以指定的时间间隔重复的向通道C发送时间值
	boom:=time.After(5e8) // 5e8 = 500000000 = 500ms 在指定的时间间隔后向通道C发送一个时间值
	for {
		select { // select 语句用于从不同的并发执行的协程中获取值
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
