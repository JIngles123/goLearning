package main

import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

// 从控制台读取输入
func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName) // 从标准输入读取输入, 并赋值给firstName和lastName, 直到碰到换行符结束
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)// 从input字符串中按照format格式读取数据, 并赋值给f, i, s
	fmt.Println("From the string we read: ", f, i, s)
}
