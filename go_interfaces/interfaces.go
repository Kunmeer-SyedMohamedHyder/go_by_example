package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area is: ", g.area())
	fmt.Println("Perimeter is: ", g.perimeter())
}

func main() {
	r := rect{length: 10, breadth: 8}
	c := circle{radius: 3}

	measure(r)
	measure(c)
}
