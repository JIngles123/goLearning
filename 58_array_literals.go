package main

import "fmt"

func main() {
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}
	var arrKeyValues = [5]string{3:"Chris", 4:"Ron"}// 只有索引3和4被赋予了值，其他元素被自动初始化为""
	// var arrKeyValue = []string{3:"Chris", 4:"Ron"}

	for i:=0;i<len(arrKeyValues);i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValues[i])
	}
}

// 输出
// Person at 0 is 
// Person at 1 is 
// Person at 2 is 
// Person at 3 is Chris
// Person at 4 is Ron