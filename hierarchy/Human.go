package hierarchy

import (
	"fmt"
)

type human struct {
	age int
}

func (h human) GetAge() int {
	return h.age
}

func (h *human) SetAge(age int) {
	h.age = age
}

func (h human) Summarise() {
	fmt.Printf("Age of human is: %v\n", h.age)
}

func NewHuman(age int) *human {
	return &human{age}
}
