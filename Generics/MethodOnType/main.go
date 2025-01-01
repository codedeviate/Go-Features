package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type NumberList[T constraints.Ordered] struct {
	values []T
}

// NewNumberList is a generic constructor
func NewNumberList[T constraints.Ordered](vals ...T) NumberList[T] {
	return NumberList[T]{values: vals}
}

// Max is a method that returns the largest value
func (nl NumberList[T]) Max() T {
	if len(nl.values) == 0 {
		var zero T
		return zero
	}
	maxVal := nl.values[0]
	for _, v := range nl.values[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	intList := NewNumberList(1, 5, 3, 9, 2)
	fmt.Println(intList.Max()) // 9

	floatList := NewNumberList(3.1, 2.4, 5.9)
	fmt.Println(floatList.Max()) // 5.9
}
