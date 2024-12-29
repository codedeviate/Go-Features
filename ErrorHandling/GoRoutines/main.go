package main

import (
	"errors"
	"fmt"
)

func doTask(results chan error) {
	results <- errors.New("task failed")
}

func main() {
	results := make(chan error, 1)
	go doTask(results)

	err := <-results
	if err != nil {
		fmt.Println("Error:", err)
	}
}
