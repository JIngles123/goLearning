package main

// 通常情况下，import "包的路径或 URL 地址" 如：import "github.com/org1/pack1”
import (
	"fmt"
	. "golearning/79_pack/pack1"
)

func main() {
	var test1 string
	test1 = ReturnStr() // 当使用. 来做为包的别名时，可以不通过包名来使用其中的项目

	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", Pack1Int)
	fmt.Printf("Float from package1: %f\n", Pack1Float)
}
