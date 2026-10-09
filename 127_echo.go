package main

import (
	"flag"
	"os"
)

var NewLine=flag.Bool("n", false, "print newline") // 定义一个默认值是false的flag

const (
	Space=" "
	Newline="\n"
)

// 模拟unix的echo功能
func main() {
	flag.PrintDefaults() // 打印flag的使用帮助信息
	flag.Parse() // 扫描参数列表，并设置flag，parse()后，所有参数全部可用
	var s string=""
	for i:=0;i<flag.NArg();i++ { // flag.NArg() 返回参数的数量
		if i>0 {
			s+=" "
			if *NewLine {
				s+=Newline
			}
		}
		s+=flag.Arg(i) // 表示第i个参数，第0个就是第1个参数，而非程序名字
	}
	os.Stdout.WriteString(s) // 将字符串写入标准输出
}
