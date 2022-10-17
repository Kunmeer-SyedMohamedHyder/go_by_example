package main

import "fmt"

func main() {

	n := 9
	if n%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	n = 40

	if n < 40 {
		fmt.Println("Fail")
	} else if n == 40 {
		fmt.Println("Just pass")
	} else {
		fmt.Println("Pass")
	}

}
