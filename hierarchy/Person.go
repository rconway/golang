package hierarchy

import (
	"fmt"
)

type Person struct {
	Human
	Name string
}

func (p Person) sayName() {
	fmt.Printf("Name of person is: %v\n", p.Name)
}

func (p Person) Summarise() {
	p.Human.Summarise()
	p.sayName()
}
