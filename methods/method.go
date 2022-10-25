package main

import "fmt"

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func (r *rect) change_height(h int) {
	r.height = h
}

func main() {
	r := rect{height: 10, width: 5}
	fmt.Println("Rectangle area is: ", r.area())
	fmt.Println("Rectangle perimeter is: ", r.perimeter())

	r.change_height(1)
	fmt.Println("Rectangle area is: ", r.area())
	fmt.Println("Rectangle perimeter is: ", r.perimeter())
}
