package main

import (
	"fmt"
	"struct_pack/struct_pack" // 模块名+文件夹名
)
// (cd 项目根目录， go mod init 模块名， go mod tidy, go run .）

func main() {
	struct1:=new(structpack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 15.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}
