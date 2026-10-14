package main

import (
	"fmt"
)

func generate_149() chan int {
	ch:=make(chan int)
	go func() {
		for i:=2;;i++ {
			ch<-i
		}
	}()
	return ch
}

func filter_149(in chan int, prime int) chan int {
	out:=make(chan int)
	go func() {
		for {
			if i:=<-in; i%prime!=0 { // 不断接收到generate生成的数，能被整除的忽略，不能被整除的写入out通道
				out<-i
			}
		}
	}()
	return out
}

func sieve_149() chan int {
	out:=make(chan int)
	go func() {
		ch:=generate_149() // ch先给出的是2
		for {
			prime:=<-ch // 执行无数遍
			ch = filter_149(ch, prime) // ch 改指向新通道filter的输出通道：后续只收到不能被该素数整除的数（合数在 filter 里被丢弃，不会传到这里）
			out<-prime // 将素数写入out通道，发回给main，35行之后每次循环读到的都是下一个素数
		}
	}()
	return out
}
// 以上3个函数都是工厂，都创建通道并返回，且使用了协程的lambda函数
// 在 Go 里，go 关键字就是用来启动一个新的 goroutine（协程）。
func main() {
	primes:=sieve_149()
	for {
		fmt.Println(<-primes)
	}
}
