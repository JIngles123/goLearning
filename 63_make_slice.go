package main
import "fmt"

func main() {
	// func make([]T, len, cap)，其中 cap 是可选参数，如果不传，则 cap = len
	var slice1 []int = make([]int, 10)
	for i:=0; i<len(slice1); i++ {
		slice1[i] = 5*i
	}

	for i:=0; i<len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i:=0; i<len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

// make([]int, 50, 100) 等同于 new([100]int)[0:50]
// new () 和 make () 的区别:
// 二者都在堆上分配内存
// new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于 &T{}。
// make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel。
// 换言之，new 函数分配内存，make 函数初始化