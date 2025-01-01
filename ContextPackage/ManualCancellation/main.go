package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    // Create a base context
    ctx := context.Background()

    // Create a cancellable context
    ctx, cancel := context.WithCancel(ctx)

    // Start a goroutine that does some work
    go func() {
        // Use a loop to mimic some ongoing work
        for {
            select {
            case <-ctx.Done():
                // Context was canceled
                fmt.Println("Worker: context canceled, stopping...")
                return
            default:
                fmt.Println("Worker: working...")
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()

    // Let the worker run for 2 seconds
    time.Sleep(2 * time.Second)

    // Cancel the context
    fmt.Println("Main: canceling context...")
    cancel()

    // Wait a bit to see the goroutine exit
    time.Sleep(1 * time.Second)
    fmt.Println("Main: done")
}
