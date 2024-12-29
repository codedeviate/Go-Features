package main

import (
	"errors"
	"fmt"
)

// Define custom error types
var ErrNotFound = errors.New("resource not found")

type ValidationError struct {
	Field string
	Msg   string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation error: field '%s' - %s", v.Field, v.Msg)
}

// Simulate a function that returns wrapped errors
func getResource(id int) error {
	if id == 0 {
		return fmt.Errorf("getResource failed: %w", ErrNotFound)
	}
	if id < 0 {
		return &ValidationError{Field: "ID", Msg: "must be a positive number"}
	}
	return nil
}

func main() {
	// Test error handling
	ids := []int{0, -1, 42}
	for _, id := range ids {
		err := getResource(id)
		if err != nil {
			// Check if the error is ErrNotFound using errors.Is
			if errors.Is(err, ErrNotFound) {
				fmt.Printf("ID %d: Error: Resource not found\n", id)
				continue
			}

			// Check if the error is of type ValidationError using errors.As
			var validationErr *ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("ID %d: Validation Error: %s\n", id, validationErr)
				continue
			}

			// Generic error handling
			fmt.Printf("ID %d: Unexpected error: %v\n", id, err)
		} else {
			fmt.Printf("ID %d: Resource retrieved successfully\n", id)
		}
	}
}
