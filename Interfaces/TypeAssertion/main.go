package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Type assertion to string
	str, ok := i.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Not a string")
	}
}
