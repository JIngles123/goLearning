package main

import (
	"fmt"
	// 用 go modeules 路径导入，不要用相对路径
	// Go 会在当前 go.mod 所在目录里解析 golearning/5_init/trans，对应到旁边的 trans/ 包。
	"golearning/5_init/trans"
)

var twoPi = 2 * trans.Pi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi)
}

// 执行命令：
// $ cd /home/jlm/Applications/Investigation/goLearning/5_init
// $ go run .