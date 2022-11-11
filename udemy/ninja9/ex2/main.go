// Package 2
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() string {
	return fmt.Sprintf("I am %v and I am %v years old", p.name, p.age)
}

func (p *person) speak2() string {
	return fmt.Sprintf("I am %v and I am %v years old", p.name, p.age)
}

type human interface {
	speak() string
}

func say_something(h human) string {
	return h.(*person).speak2()
}

func main() {
	p := person{
		name: "Syed",
		age:  22,
	}
	s := p.speak()
	fmt.Println(s)
	s = say_something(&p)
	fmt.Println(s)

}
