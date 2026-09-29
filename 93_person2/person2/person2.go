package person2

type Person struct {
	firstName string
	lastName string
}

// 外部无法直接访问firstName，只能通过FirstName()方法访问

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}
