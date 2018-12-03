package hierarchy

import (
	"fmt"
)

type person struct {
	human
	name string
}

func (p person) GetName() string {
	return p.name
}

func (p *person) SetName(name string) {
	p.name = name
}

func (p person) Summarise() {
	p.human.Summarise()
	fmt.Printf("Name of person is: %v\n", p.name)
}

func NewPerson(age int, name string) *person {
	//return &person{ human{age}, name }
	return &person{*NewHuman(age), name}
}
