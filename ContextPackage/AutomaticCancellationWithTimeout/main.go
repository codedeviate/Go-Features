package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    // Create a context with a 2-second timeout
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Always call cancel to free resources

    go func() {
        <-ctx.Done()
        fmt.Println("Goroutine: context done:", ctx.Err())
    }()

    // Simulate some work that takes 3 seconds
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Main: finished work without cancellation")
    case <-ctx.Done():
        fmt.Println("Main: context canceled:", ctx.Err())
    }
}
