package main

import (
	"fmt"
	"math"
)

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements the Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implements the Shape interface
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 4}

	printShapeDetails(circle)    // Circle satisfies Shape
	printShapeDetails(rectangle) // Rectangle satisfies Shape
}
