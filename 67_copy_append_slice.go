package main

import "fmt"

func main() {
	sl_from := []int{1,2,3}
	sl_to := make([]int,10)

	// 如果想增加切片的容量，必须创建一个新的更大的切片并把原分片的内容都拷贝过来
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1,2,3}
	sl3 = append(sl3,4,5,6)// 追加的元素必须和原切片的元素同类型
	fmt.Println(sl3)
}

