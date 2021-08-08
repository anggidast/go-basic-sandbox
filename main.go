package main

import (
	"errors"
	"fmt"
	"math"
)

// shape interface
type shape interface { // * group types together, base on their methods
	area() float64
	circumf() float64
}

type square struct { // * square type have area and circumf method, so its shape interface
	length float64
}
type circle struct { // * circle type have area and circumf method, so its shape interface
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

// ? ERROR INTERFACE (HANDLER)
// * Go has built-in package interface for error handler for representing an errors condition, with the nil value representing no error
// * import package errors
// * remember that error is one of data type in Go

func Devide(value int, devider int) (int, error) {
	if devider == 0 {
		return 0, errors.New("Devided by zero")
	} else {
		return value / devider, nil
	}
}

// func main() {
// 	result, err := Devide(8, 0)
// 	if err == nil {
// 		fmt.Println("Result:", result)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }
