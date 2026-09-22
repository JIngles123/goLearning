package main
import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // index[2,5), 不使用指针指向slice，切片本身已经是一个引用类型，即本身就是一个指针

	for i:=0; i<len(arr1); i++ {
		arr1[i] = i
	}

	for i:=0; i<len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// cap 不是「这段切片自己的长度」，而是「从切片的起始位置，到底层数组末尾，还有多少元素」
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	// 切片的索引是相对于切片自己的起点，不是相对于底层数组的起点
	slice1 = slice1[0:4] // [0,4)
	for i:=0; i<len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// slice1 = slice1[0:7]
	// panic: runtime error: slice bounds out of range [0:7] with capacity 4
}

// 输出：
// Slice at 0 is 2
// Slice at 1 is 3
// Slice at 2 is 4
// The length of arr1 is 6
// The length of slice1 is 3
// The capacity of slice1 is 4
// Slice at 0 is 2
// Slice at 1 is 3
// Slice at 2 is 4
// Slice at 3 is 5
// The length of slice1 is 4
// The capacity of slice1 is 4