package main

import "fmt"

// Stack is a LIFO data structure for elements of type T.
type Stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

// Pop removes the top element of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.data) - 1
	elem := s.data[index]
	s.data = s.data[:index]
	return elem, true
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	val, ok := intStack.Pop()
	fmt.Println(val, ok) // 20 true

	stringStack := Stack[string]{}
	stringStack.Push("go")
	stringStack.Push("generics")
	v, ok := stringStack.Pop()
	fmt.Println(v, ok) // generics true
}
