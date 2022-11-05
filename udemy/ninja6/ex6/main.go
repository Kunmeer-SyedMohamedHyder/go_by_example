package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

type circle struct {
	radius int
}

func (s square) area() float64 {
	return float64(s.side * s.side)
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("Area of the shape is: %v\n", s.area())
}

func main() {
	sq := square{
		side: 4,
	}
	c := circle{
		radius: 2,
	}
	info(sq)
	info(c)

}
