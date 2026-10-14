package main

// 打印素数，使用选择器作为算法，每个prime都有一个选择器

import "fmt"

func generate(ch chan int) {
	for i:=2;;i++ {
		ch<-i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i:=<-in
		if i%prime!=0 {
			out<-i
		}
	}
}

func main() {
	ch:=make(chan int)
	go generate(ch)
	for {
		prime:=<-ch
		fmt.Print(prime, " ")
		ch1:=make(chan int)
		go filter(ch, ch1, prime) // 拷贝整数到输出通道，丢弃能被prime整除的整数，每个prime开一个新协程，生成器和选择器
		ch=ch1
	}
}
