package main

import "os"

func main() {
	os.Stdout.WriteString("hello, world\n") // 输出到屏幕
	f,_:=os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0) // 只写模式创建打开文件test
	defer f.Close()
	f.WriteString("hello, world in a file\n") // 不使用缓冲区，直接将内容写入文件
}
