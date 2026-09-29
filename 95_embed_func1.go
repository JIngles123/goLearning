package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

func main() {
	c:=new(Customer)
	c.Name="Barak Obama"
	c.log=new(Log)
	c.log.msg="1 - Yes we can!"

	// shorter
	c=&Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg+="\n"+s
}

// 这里是被 Println 自动调用，有这个函数则打印string，没有则打印引用
func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
