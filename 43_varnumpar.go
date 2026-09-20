package main

import "fmt"

func main() {
	x:=min(1,3,2,0)
	fmt.Printf("The minimum is: %d\n", x)

	// []int{...} 是切片字面量（slice literal），创建一个元素类型为 int 的切片
	slice:=[]int{7,9,3,5,1}
	// slice... 是切片展开操作（slice expansion），将切片中的所有元素作为参数传递给 min 函数
	x=min(slice...)
	fmt.Printf("The minimum of slide is: %d\n", x)
}

// x ...int 表示这是一个可变参数函数，可以接收任意个 int
func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min:=s[0]
	for _,v:=range s {
		if v<min {
			min = v
		}
	}
	return min
}
