package main

import (
	"fmt"
	"bufio"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

// 使用bufio包中的缓冲读取（buffered reader）来读取数据
func main() {
	inputReader = bufio.NewReader(os.Stdin) // 创建一个读取器，将从指定读取器（os.Stdin）读取内容
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
