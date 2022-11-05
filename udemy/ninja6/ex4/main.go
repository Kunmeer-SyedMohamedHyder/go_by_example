package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("I am %v and am %v years old..", p.last, p.age)
}

func main() {
	p := person{
		first: "Syed",
		last:  "Hyder",
		age:   22,
	}
	p.speak()
}
