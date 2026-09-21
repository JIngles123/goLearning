package main

import "fmt"

func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	for i:=0;i<3;i++ {
		fp(&[3]int{i, i*i, i*i*i})
	}
}

// 输出
// &[0 0 0]
// &[1 1 1]
// &[2 4 8]

