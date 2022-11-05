package main

import "fmt"

func sum1(arr ...int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func sum2(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func modify_arr(arr []int) {
	arr[1] = 100
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ans1 := sum1(arr...)

	modify_arr(arr)
	ans2 := sum2(arr)

	fmt.Println("Sum1 of the array is : ", ans1)
	fmt.Println("Sum2 of the array is : ", ans2)

	fmt.Println(arr)

}
