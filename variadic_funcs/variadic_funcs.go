package main

import "fmt"

func demo(nums ...int) {
	fmt.Print(nums)

	ans := 0
	for _, num := range nums {
		ans += num
	}

	fmt.Println(ans)
}

func main() {
	demo(1, 2, 3, 4)

	s := []int{5, 6, 7, 8, 9}
	demo(s...)

	demo(s[1:]...)
}
