package main

import "fmt"

type Person struct {
	Name string
}

// Pointer receiver
func (p *Person) SetName(newName string) {
	p.Name = newName
}

type Employee struct {
	Person
	ID int
}

func main() {
	// e is a value
	e := Employee{
		Person: Person{Name: "Charlie"},
		ID:     42,
	}

	// Even though e is not a pointer, we can still call SetName,
	// and Go will take the address of e.Person automatically.
	e.SetName("Charlotte")
	fmt.Println("Employee Name:", e.Name)

	// If we have a pointer, it works similarly
	ep := &Employee{
		Person: Person{Name: "Dave"},
		ID:     100,
	}
	ep.SetName("David")
	fmt.Println("Employee Pointer Name:", ep.Name)
}
