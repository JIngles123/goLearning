package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string

	// 打印整数类型的大小（多少位的系统）
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	
	// 将字符串转换为整数
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)

	// 将整数转换为字符串
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}