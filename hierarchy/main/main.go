package main

import (
	"fmt"
	"github.com/rconway/golang/hierarchy"
)

func main() {
	// Test human
	fmt.Println("---")
	h := hierarchy.Human{12}
	h.Summarise()

	// Test human again
	fmt.Println("---")
	h = hierarchy.Human{34}
	h.Summarise()

	// Test person
	fmt.Println("---")
	p := hierarchy.Person{hierarchy.Human{56}, "fredbob"}
	p.Summarise()

}
