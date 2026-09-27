package main

import (
	"fmt"
	"strings"
)

// go中，结构体和它所包含的数据在内存中是以连续块的形式存在的

type Person struct {
	firstNme string 
	lastName string
}

func upPerson(p *Person) {
	p.firstNme = strings.ToUpper(p.firstNme)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1 var -> a struct with named fields
	var pers1 Person
	pers1.firstNme = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstNme, pers1.lastName)

	// 2 new-> a pointer to a new struct
	pers2:=new(Person)
	pers2.firstNme = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstNme, pers2.lastName)

	// 3 struct literal
	pers3:=&Person{"Chris","Woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstNme, pers3.lastName)

	pers4:=&Person{firstNme: "Chris", lastName: "Woodward"}
	upPerson(pers4)
	fmt.Printf("The name of the person is %s %s\n", pers4.firstNme, pers4.lastName)
}
