package main

import "fmt"

func main() {
	str := ""
	for i := 0; i < 10; i++ {
		str += "A" // Creates a new string each iteration
	}
	fmt.Println(str)
}
