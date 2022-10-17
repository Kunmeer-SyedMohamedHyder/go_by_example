package main

import "fmt"

func plus(a, b int) int {

	c := a + b
	return c
}

func plusplus(a, b, c int) int {
	d := a + b + c
	return d
}

func main() {

	a, b, c := 10, 20, 30
	ans := plus(a, b)
	fmt.Println(ans)

	ans = plusplus(a, b, c)
	fmt.Println(ans)

}
