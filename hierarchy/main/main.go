package main

import (
	"fmt"
	"github.com/rconway/golang/hierarchy"
)

func main() {
	// Test human
	fmt.Println("---")
	h := hierarchy.NewHuman(12)
	h.Summarise()

	// Test human again
	fmt.Println("---")
	h = hierarchy.NewHuman(34)
	h.Summarise()

	// Test person
	fmt.Println("---")
	p := hierarchy.NewPerson(56, "fredbob")
	p.Summarise()

	// Modify person
	fmt.Println("---")
	p.SetAge(123)
	p.Summarise()
	p.SetName("barry")
	p.Summarise()

	// Test student
	fmt.Println("---")
	s := hierarchy.NewStudent(11, "james dean", "raglan")
	s.Summarise()

	// Modify student
	s.SetAge(42)
	s.SetName("Stephen Hawking")
	s.SetSchool("Cambridge")
	s.Summarise()
}
