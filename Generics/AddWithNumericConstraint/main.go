package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Numeric is a custom constraint satisfied by integer and float types.
type Numeric interface {
	constraints.Integer | constraints.Float
}

// Add sums two numeric values.
func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add[int](2, 3))           // 5
	fmt.Println(Add[float64](3.14, 2.71)) // 5.85
}
