package main

import "fmt"

type Person struct {
	Name string
}

type Employee struct {
	Person
	Name string // Shadowing Person.Name
}

func main() {
	e := Employee{
		Person: Person{Name: "Eve Person"},
		Name:   "Eve Employee",
	}

	// Accessing the shadowed fields
	fmt.Println("Employee Name:", e.Name)               // "Eve Employee"
	fmt.Println("Embedded Person Name:", e.Person.Name) // "Eve Person"
}
