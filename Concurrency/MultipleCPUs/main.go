package main

import (
	"fmt"
	"runtime"
	"sync"
)

func task(id int) {
    fmt.Printf("Task %d is running\n", id)
}

func main() {
    // Determine and set the maximum number of processors (CPUs) that Go should use
    numCPUs := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPUs)

    var wg sync.WaitGroup

    // Launch multiple goroutines
    for i := 1; i <= 10; i++ {
        wg.Add(1) // Increment the WaitGroup counter
        go func(id int) {
            defer wg.Done() // Decrement when the goroutine completes
            task(id)
        }(i)
    }

    // Wait for all goroutines to complete
    wg.Wait()

    fmt.Println("All tasks completed.")
}
