package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Multiply[T constraints.Integer | constraints.Float](a, b T) T {
	return a * b
}

func main() {
	// Type inference from arguments; no need to explicitly specify [int] or [float64].
	fmt.Println(Multiply(3, 4))     // infers T = int
	fmt.Println(Multiply(2.5, 1.5)) // infers T = float64
}
