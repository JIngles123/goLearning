package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm Marc, Hi."

	fmt.Printf("The position of \"Marc\" is: %d\n", strings.Index(str, "Marc"))

	// 第一次出现的位置
	fmt.Printf("The position of the first instance of \"Hi\" is: %d\n", strings.Index(str, "Hi"))

	// 最后一次出现的位置
	fmt.Printf("The position of the last instance of \"Hi\" is: %d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: %d\n", strings.Index(str, "Burger"))
}
