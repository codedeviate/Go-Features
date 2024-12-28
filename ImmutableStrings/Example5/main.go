package main

import "fmt"

func main() {
	str1 := "Immutable"
	str2 := str1

	fmt.Println("Before modification:")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	// "Modify" str2
	str2 = "New String"
	fmt.Println("\nAfter modification:")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
}
