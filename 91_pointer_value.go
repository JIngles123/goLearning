package main

import (
	"fmt"
)

type B struct {
	thing int
}

/*
说明：
change() 接受一个指向 B 的指针，并改变它内部的成员；
write() 接受通过拷贝接受 B 的值并只输出 B 的内容。

可以正常编译，但是开始的 b 没有被改变。

对于类型 T，如果在 *T 上存在方法 Meth()，并且 t 是这个类型的变量，
那么 t.Meth() 会被自动转换为 (&t).Meth()。

=> 指针方法和值方法都可以在指针或非指针上被调用
*/

func (b *B) change() {
	b.thing=1
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	// b1是值
	var b1 B
	b1.change()
	fmt.Println(b1.write())

	// b2是指针
	b2:=new(B)
	b2.change()
	fmt.Println(b2.write())
}
