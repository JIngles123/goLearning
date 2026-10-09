package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who:="Alice "
	if len(os.Args)>1 { // 用于在程序启动后读取命令输入的参数，切片，从索引1开始，索引0是程序名
		who+=strings.Join(os.Args[1:], " ")// Join是以空格为间隔连接这些参数
	}
	fmt.Println("Good Morning", who)
}

// 在终端运行：go run 126_os_args.go Alice Bob Charlie
// 输出：Good Morning Alice Alice Bob Charlie
