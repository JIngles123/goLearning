package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// 结合使用缓冲读取文件和命令行 flag。
func cat(r *bufio.Reader) {
	for {
		buf, err:=r.ReadBytes('\n')
		if err==io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf) // 将buf写入标准输出
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg()==0 {
        // 如果不加参数，则输入什么屏幕就打印什么。
		cat(bufio.NewReader(os.Stdin))
	}

	for i:=0;i<flag.NArg();i++ {
		f, err:=os.Open(flag.Arg(i)) // 打开文件
		if err!=nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f)) // 用buffer读取文件内容并打印到标准输出
	}
}
