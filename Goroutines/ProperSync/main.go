package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup, num int) {
	defer wg.Done() // Mark as done
	fmt.Println(num)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter
		go printNumber(&wg, i)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines completed")
}
