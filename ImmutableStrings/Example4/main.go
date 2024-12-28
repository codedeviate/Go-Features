package main

import "fmt"

func main() {
	str := "Go语言" // UTF-8 string
	runes := []rune(str)

	// Change the second character
	runes[2] = '码' // Replace '语' with '码'

	newStr := string(runes)
	fmt.Println("Original String:", str)
	fmt.Println("Modified String:", newStr)
}
