package main

import "fmt"

func incrementer(val int) (func(), func(i int)) {

	inc1 := func() {
		val++
		fmt.Println(val)
	}

	inc2 := func(i int) {
		val += i
		fmt.Println(val)
	}
	return inc1, inc2
}

func main() {

	inc1, inc2 := incrementer(10)
	i, _ := incrementer(1)
	inc1()
	inc1()
	inc2(10)
	i()
}
