package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	// 将字符串按空格分割成切片
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl { // 遍历切片,第一个值是索引,第二个值是切片中的元素
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	// 将字符串按指定字符分割成切片
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	// 将切片中的元素用指定字符连接成一个字符串
	str3 := strings.Join(sl2, ";")
	fmt.Printf("Joined: %s\n", str3)
}