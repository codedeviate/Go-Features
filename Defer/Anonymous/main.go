package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Deferred anonymous function executed")
	}()
	fmt.Println("Main function executed")
}
