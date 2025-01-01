package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person  // Embedded struct
	ID      int
	Company string
}

func main() {
	e := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		ID:      1001,
		Company: "TechCorp",
	}

	// Accessing fields directly
	fmt.Println("Name:", e.Name) // promoted from Person
	fmt.Println("Age:", e.Age)   // promoted from Person
	fmt.Println("ID:", e.ID)
	fmt.Println("Company:", e.Company)
}
