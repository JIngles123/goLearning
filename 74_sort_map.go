package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
							"delta": 87, "echo": 56, "foxtrot": 12,
							"golf": 34, "hotel": 16, "indio": 87,
							"juliet": 65, "kili": 43, "lima": 98}
)

/*
map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序。

如果要为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包），
然后可以使用切片的 for-range 方法打印出所有的 key 和 value
*/

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v /", k,v)
	}

	keys:=make([]string, len(barVal))
	i:=0
	for k,_:=range barVal {// key自身是索引，对k的排序也是对map的排序
		keys[i]=k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")

	for _, k := range keys { // var keys []string
		fmt.Printf("Key: %v, Value: %v /", k, barVal[k])
	}
}
