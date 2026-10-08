package main

import (
	"fmt"
)

type List []int

type Appender interface {
	Append(int)
}

func (l *List) Append(val int) { // 接收者是指针的方法
	*l=append(*l, val)
}

func CountInto(a Appender, start, end int) {
	for i:=start; i<=end; i++ {
		a.Append(i) // 在接口中调用方法
	}
}

// ===============================
type Lener interface {
	Len() int
}

func (l List) Len() int { // 接收者是值的方法
	return len(l)
}

func LongEnough(l Lener) bool {
	return l.Len()*10>42 // 在接口中调用方法
}

func main() {
	var lst List
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst:=new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}

/*
methodset1.go 中我们看到，作用于变量上的方法实际上是不区分变量到底是指针还是值的。
当碰到接口类型值时，这会变得有点复杂，原因是接口变量中存储的具体值是不可寻址的，幸运的是，
如果使用不当编译器会给出错误。

总结：
在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的：
1、指针方法可以通过指针调用
2、值方法可以通过值调用
3、接收者是值的方法可以通过指针调用，因为指针会首先被解引用
4、接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
5、将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都可以在此值上被调用，因此不正确的赋值在编译期就会失败。

*/
