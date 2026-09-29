package main

import (
	"person2/person2"
	"fmt"
)

func main() {
	p:=new(person2.Person)
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}
