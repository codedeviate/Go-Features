package main

import "fmt"

func main() {
	var i interface{}
	fmt.Println(i == nil) // true

	var p *int
	i = p
	fmt.Println(i == nil) // false
}
