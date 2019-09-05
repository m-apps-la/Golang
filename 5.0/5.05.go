package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func main() {
	s := square{
		length: 15,
	}
	c := circle{
		radius: 12.345,
	}
	info(s)
	info(c)
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}
