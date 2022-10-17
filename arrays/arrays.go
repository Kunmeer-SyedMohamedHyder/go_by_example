package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Get: ", a)

	a[4] = 10
	fmt.Println("Set: ", a)

	var twoD [2][3]int

	// var b [3]int

	for i := 0; i < len(twoD); i++ {
		// b = twoD[i]
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D Array: ", twoD)
}
