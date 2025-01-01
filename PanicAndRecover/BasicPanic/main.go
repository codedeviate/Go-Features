package main

import "fmt"

func main() {
	fmt.Println("Before the panic")
	panic("Something went wrong!")
	fmt.Println("After the panic") // This will never run
}
