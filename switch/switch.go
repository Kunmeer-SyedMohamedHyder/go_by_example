package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10

	fmt.Print("2 is spelled as ")
	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	t := time.Now()

	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Enjoy your weekend")
	default:
		fmt.Println("Sadd Weekdays")
	}

	switch {
	case t.Hour() < 12:
		fmt.Println("It's berfore noon")
	default:
		fmt.Println("It's after noon")
	}

	whoAmI := func(i interface{}) {
		switch tp := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Printf("I am special type : %T\n", tp)

		}
	}

	whoAmI(true)
	whoAmI(n)
	whoAmI(t)
}
