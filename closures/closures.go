package main

import "fmt"

func countdown(start, stop int) func() {

	return func() {
		if start == stop {
			fmt.Println("Already launched !")
		} else {
			fmt.Println(start)
			start--
		}
	}

}

// func main() {
// 	launch := countdown(10, 0)
// 	for i := 0; i < 12; i++ {
// 		launch()
// 	}
// }
