package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Printf("Area: %v\n", s.getArea())
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{base: 1, height: 2}
	s := square{sideLength: 2}

	printArea(t)
	printArea(s)
}
