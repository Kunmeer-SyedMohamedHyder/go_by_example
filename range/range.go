package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5}

	for ind, val := range a {
		fmt.Println(ind, val)
	}

	sum := 0
	for _, val := range a {
		sum += val
	}
	fmt.Println(sum)

	for key, val := range map[string]int{"hello": 1, "world": 2} {
		fmt.Println(key, "-------->", val)
	}

	for _, bt := range "hello" {
		fmt.Printf("%c\n", bt)
	}
}
