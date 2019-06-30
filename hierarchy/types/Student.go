package types

import (
	"fmt"
)

type student struct {
	person
	school string
}

func (s student) GetSchool() string {
	return s.school
}

func (s *student) SetSchool(school string) {
	s.school = school
}

func (s student) Summarise() {
	s.person.Summarise()
	fmt.Printf("Student attends school: %v\n", s.school)
}

func NewStudent(age int, name string, school string) *student {
	return &student{*NewPerson(age, name), school}
}
