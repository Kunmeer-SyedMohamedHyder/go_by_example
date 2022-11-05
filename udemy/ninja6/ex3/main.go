package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("anonymous here....")
	}()

	fmt.Println("Foo called..")
}

func bar() {
	fmt.Println("Bar called..")
}
