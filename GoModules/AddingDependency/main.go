package main

import (
	"fmt"

	"github.com/google/uuid" // external dependency
)

func main() {
    id := uuid.New()
    fmt.Println("Generated UUID:", id)
}
