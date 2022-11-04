package main

import "fmt"

type hello int

var x hello

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 56
	fmt.Println(x)
}
