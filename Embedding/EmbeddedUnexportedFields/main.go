package main

import "fmt"

type internalData struct {
	secret string
}

// Exported type embedding an unexported field
type PublicType struct {
	internalData
}

func NewPublicType(secret string) *PublicType {
	return &PublicType{internalData{secret}}
}

func main() {
	p := NewPublicType("hidden value")
	// We can't access p.secret directly because it's inside unexported internalData
	// But we can expose methods on PublicType if we want to manipulate the data.
	fmt.Println(p) // &{{hidden value}}
}
