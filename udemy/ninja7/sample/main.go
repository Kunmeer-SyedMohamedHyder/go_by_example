package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	people := ByAge{
		{"Hyder", 22},
		{"Goks", 21},
		{"Rahima", 18},
	}

	fmt.Printf("Before sort : %v\n", people)

	sort.Sort(people)
	fmt.Printf("After sort : %v\n", people)

	people2 := ByName(people)

	fmt.Printf("Before sort : %v\n", people2)

	sort.Sort(people2)
	fmt.Printf("After sort : %v\n", people2)
}
