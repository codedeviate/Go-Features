package main

import "fmt"

func example() (result int) {
	defer func() {
		result += 5
	}()
	return 10
}

func main() {
	fmt.Println("Result:", example())
}
