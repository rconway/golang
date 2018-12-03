package hierarchy

import (
	"fmt"
)

type Human struct {
	Age int
}

func (h Human) sayAge() {
	fmt.Printf("Age of human is: %v\n", h.Age)
}

func (h Human) Summarise() {
	h.sayAge()
}
