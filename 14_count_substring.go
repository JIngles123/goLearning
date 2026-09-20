package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	// 统计字符串中指定子串出现的次数
	fmt.Printf("Number of H's in %s is: %d\n", str, strings.Count(str, "H"))
	fmt.Printf("Number of g's in %s is: %d\n", manyG, strings.Count(manyG, "g"))
}