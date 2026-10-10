package main

import (
	"fmt"
	// "os/exec"
	"os"
)

/*
os 包有一个 StartProcess 函数可以 调用或启动外部系统命令和二进制可执行文件；
它的第一个参数是要运行的进程，
第二个参数用来传递选项或参数，
第三个参数是含有系统环境基本信息的结构体。

这个函数返回被启动进程的 id（pid），或者启动失败返回错误。
*/
func main() {
	env:=os.Environ()
	procAttr:=&os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	// 1st example: list files
	pid, err:=os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	
	// 2nd example: show all processes
	// pid, err := os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, procAttr)

	if err!=nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
}
