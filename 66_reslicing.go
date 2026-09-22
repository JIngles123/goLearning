package main
import "fmt"

func main() {
	// slice1 := make([]type, start_length, capacity) 创建一个长度为 start_length，容量为 capacity 的切片
	slice1 := make([]int, 0, 10)

	// 切片可以反复扩展直到占据整个相关数组。
	for i:=0; i<cap(slice1); i++ {
		slice1 = slice1[0:i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	for i:=0; i<len(slice1); i++ {
		fmt.Printf("Slice at %d is: %d\n", i, slice1[i])
	}
}
