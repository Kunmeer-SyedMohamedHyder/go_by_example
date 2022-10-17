package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("-------------------")

	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}
	fmt.Println("-------------------")

	for {
		fmt.Println("Loooppp")
		break
	}
	fmt.Println("-------------------")

	for n := 0; n < 5; n++ {
		if n == 2 {
			fmt.Println("Loop")
			continue
		}
		fmt.Println(n)
	}
}
