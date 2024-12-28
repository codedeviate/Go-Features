package main

import "fmt"

func describeType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("String:", v)
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	describeType(42)
	describeType("hello")
	describeType(3.14)
}
