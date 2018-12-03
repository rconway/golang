package hierarchy

import (
	"fmt"
)

type human struct {
	age int
}

func (h *human) SetAge(age int) {
    h.age = age
}

func (h human) sayAge() {
	fmt.Printf("Age of human is: %v\n", h.age)
}

func (h human) Summarise() {
	h.sayAge()
}

func NewHuman(age int) *human {
    return &human{age}
}
