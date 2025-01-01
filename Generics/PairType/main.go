package main

import "fmt"

// Pair holds two values of possibly different types.
type Pair[A, B any] struct {
	First  A
	Second B
}

func main() {
	intString := Pair[int, string]{First: 10, Second: "hello"}
	fmt.Println(intString.First, intString.Second) // 10 hello

	stringFloat := Pair[string, float64]{First: "pi", Second: 3.14159}
	fmt.Println(stringFloat.First, stringFloat.Second) // pi 3.14159
}
