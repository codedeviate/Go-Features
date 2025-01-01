package main

import "fmt"

// Stringer is a standard interface from fmt.
type Stringer interface {
	String() string
}

// PrintStrings prints the string representation of any type that implements Stringer.
func PrintStrings[T Stringer](items []T) {
	for _, item := range items {
		fmt.Println(item.String())
	}
}

// Example type implementing Stringer
type User struct {
	Name string
}

func (u User) String() string {
	return "User: " + u.Name
}

func main() {
	users := []User{{Name: "Alice"}, {Name: "Bob"}}
	PrintStrings(users)
}
