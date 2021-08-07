package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface { // * group types together, base on their methods
	area() float64
	circumf() float64
}

type square struct { // * square have area and circumf method, so its interface
	length float64
}
type circle struct { // * circle have area and circumf method, so its interface
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) { // * shape is interfaces
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
