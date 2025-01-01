package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Person struct {
	Name string
	Address
}

type Employee struct {
	Person
	ID int
}

func main() {
	e := Employee{
		Person: Person{
			Name: "Frank",
			Address: Address{
				City:  "Metropolis",
				State: "NY",
			},
		},
		ID: 9999,
	}

	// Access fields from nested embedding
	fmt.Println("Name:", e.Name)
	fmt.Println("City:", e.City)   // From Address
	fmt.Println("State:", e.State) // From Address
	fmt.Println("ID:", e.ID)
}
