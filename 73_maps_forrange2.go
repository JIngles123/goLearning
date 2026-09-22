package main

import "fmt"

func main() {
	items:=make([]map[int]int,5)
	for i:=range items {// 通过索引使用切片的 map 元素
		items[i]=make(map[int]int,1)
		items[i][1]=2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	items2:=make([]map[int]int,5)
	for _, item:=range items2 {// 此项只是 map 值的一个拷贝，所以真正的 map 元素没有得到初始化
		item = make(map[int]int,1)
		item[1]=2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
	
}

