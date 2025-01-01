package main

import (
	"fmt"

	"golang.org/x/exp/constraints" // requires "go get golang.org/x/exp/constraints"
)

// Min returns the smaller of two values.
// T must be ordered (i.e., supports < or > comparisons).
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Min(3, 5))         // int
	fmt.Println(Min(4.2, 2.1))     // float64
	fmt.Println(Min("apple", "z")) // string
}
