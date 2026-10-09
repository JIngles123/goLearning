package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered") // 直接输出到标准输出

	buf:=bufio.NewWriter(os.Stdout) // 创建一个缓冲写入器，将输出写入到缓冲区；NewWriter适合任何形式的缓冲写入
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered") // 将输出写入到缓冲区
	buf.Flush() // 在缓冲写入的最后一定要使用Flush(),否则最后的输出不会被写入到标准输出，注释掉后，输出不会被写入到标准输出，因为缓冲区没有被刷新
}

// fmt.Fprintf是写入一个io.writer接口类型的变量，按照指定的格式向第一个参数内写入字符串
// 第一个参数是io.writer接口类型的变量，第二个参数是格式化字符串，第三个参数是可变参数
