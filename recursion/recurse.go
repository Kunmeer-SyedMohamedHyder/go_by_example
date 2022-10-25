package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	n := 7
	ans := factorial(n)

	fmt.Println("Factorial is: ", ans)

	var fib func(n int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		} else {
			return fib(n-1) + fib(n-2)
		}
	}

	n = 7
	ans = fib(n)

	fmt.Println("Fibonacci is: ", ans)
}
