package main

import "fmt"

// Filter returns a new slice containing all elements from 's' for which 'keep' returns true.
func Filter[T any](s []T, keep func(T) bool) []T {
	var result []T
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	isEven := func(n int) bool { return n%2 == 0 }

	evenNums := Filter(nums, isEven)
	fmt.Println(evenNums) // [2 4]

	strs := []string{"apple", "banana", "cherry"}
	hasFiveLetters := func(s string) bool { return len(s) == 5 }
	fiveLetterStrs := Filter(strs, hasFiveLetters)
	fmt.Println(fiveLetterStrs) // ["apple"]
}
