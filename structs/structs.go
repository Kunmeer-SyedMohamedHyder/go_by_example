package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {

	p := Person{name: name}
	p.age = 40

	return &p
}

func main() {
	fmt.Println(Person{name: "Hyder", age: 22})
	fmt.Println(&Person{name: "Syed", age: 23})
	fmt.Println(newPerson("Sarika"))

}
