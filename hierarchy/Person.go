package hierarchy

import (
	"fmt"
)

type person struct {
	human
	name string
}

func (p person) sayName() {
	fmt.Printf("Name of person is: %v\n", p.name)
}

func (p person) Summarise() {
	p.human.Summarise()
	p.sayName()
}

func NewPerson(age int, name string) *person {
    //return &person{ human{age}, name }
    return &person{ *NewHuman(age), name }
}
