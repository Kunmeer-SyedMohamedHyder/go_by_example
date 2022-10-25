package main

import "fmt"

func zeroval(n int) {
	fmt.Println("n value in zeroval: ", n)
	n = 0
}
func zeroptr(n *int) {
	fmt.Println("n value in zeroptr: ", n)
	*n = 0
}
func main() {
	n := 7

	fmt.Println("n value before zeroval: ", n)
	zeroval(n)
	fmt.Println("n value ater zeroval: ", n)

	fmt.Println("n value before zeroptr: ", n)
	zeroptr(&n)
	fmt.Println("n value after zeroptr: ", n)

	fmt.Println("Address of zeroptr: ", &n)
}
