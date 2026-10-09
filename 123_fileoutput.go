package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	/*
	通常会用到以下标志：
		os.O_RDONLY：只读
		os.O_WRONLY：只写
		os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
		os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0。
	*/
	outputFile, outputError:=os.OpenFile("output.dat",
        os.O_WRONLY|os.O_CREATE, 0666)// 文件名，一个或多个标志，文件权限(写文件时，固定0666)
	if outputError!=nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	/*
	如果写入的东西很简单，可使用 fmt.Fprintf(outputFile, “Some test data.\n”) 直接将内容写入文件。
	fmt 包里的 F 开头的 Print 函数可以直接写入任何 io.Writer，包括文件
	*/
	outputWriter:=bufio.NewWriter(outputFile)// 创建一个写入器(缓冲区)对象
	outputString:="hello world!\n"

	for i:=0;i<10;i++ {
		outputWriter.WriteString(outputString) // 将字符串写入缓冲区
	}
	outputWriter.Flush() // 将缓冲区的内容全部写入到文件中
}
