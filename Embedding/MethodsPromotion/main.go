package main

import "fmt"

type Person struct {
	Name string
}

// Method on Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

type Employee struct {
	Person
	Company string
}

func main() {
	e := Employee{
		Person:  Person{Name: "Bob"},
		Company: "InnoTech",
	}

	// Directly calling the method from the embedded struct
	e.Greet()
}
