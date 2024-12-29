package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("invalid configuration")
}

func initialize() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("initialization failed: %w", err)
	}
	return nil
}

func main() {
	err := initialize()
	if err != nil {
		fmt.Println("Error:", err)

		// Unwrap the error
		unwrapped := errors.Unwrap(err)
		fmt.Println("Underlying Error:", unwrapped)
	}
}
